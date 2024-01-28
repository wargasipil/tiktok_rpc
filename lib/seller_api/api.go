package seller_api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/schema"
	"github.com/pdcgo/common_conf/pdc_common"
	retry "github.com/sethvargo/go-retry"
)

type Session interface {
	GetCookies() []*http.Cookie
	SaveSession() error
	Sync() error
	Update(cookies []*http.Cookie) error
	AddToHttpRequest(req *http.Request)
	UserAgent() string
	FindCookie(string) string
}

type SellerCacheData struct {
	OecSellerID string `json:"oec_seller_id"`
}

type SellerCache interface {
	Get() (*SellerCacheData, error)
	SetAction(key string, handler func() (*SellerCacheData, error))
}

var ClientApi *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 5,
	},
	Timeout: 30 * time.Second,
}

type SellerApi struct {
	Username string
	Session  Session
	encoder  *schema.Encoder
	Data     SellerCache
}

func NewSellerApi(username string, session Session, cache SellerCache) (*SellerApi, func() error, error) {
	saveSession := func() error {
		err := session.SaveSession()
		if err != nil {
			pdc_common.ReportError(err)
		}

		return nil
	}

	account := SellerApi{
		Session: session,
		encoder: NewEncoder(),
		Data:    cache,
	}

	cache.SetAction(username, func() (*SellerCacheData, error) {
		data, err := account.CommonInfo()
		datacache := SellerCacheData{
			OecSellerID: data.Data.AccountID,
		}
		return &datacache, err
	})

	return &account, saveSession, nil
}

func (api *SellerApi) SetHeader(req *http.Request, custom map[string]string) {

	heads := map[string]string{
		"Content-Type": "application/json",
		"Origin":       "https://affiliate.tiktok.com",
		"User-Agent":   api.Session.UserAgent(),
	}

	for key, val := range heads {
		req.Header.Set(key, val)
	}

	for key, val := range custom {
		req.Header.Set(key, val)
	}
}

func (api *SellerApi) NewRequestJSON(method, ur string, query any, body any) *http.Request {
	data, err := json.Marshal(body)
	// log.Println(string(data))
	if err != nil {
		log.Println(err)
	}
	return api.NewRequest(method, ur, query, bytes.NewBuffer(data))
}

func (api *SellerApi) NewRequest(method, ur string, query any, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, ur, body)
	if err != nil {
		log.Println(err)
	}
	// setting query'
	if query != nil {
		q := req.URL.Query()
		api.encoder.Encode(query, q)
		req.URL.RawQuery = q.Encode()
		// log.Println(q.Encode())
	}
	api.Session.AddToHttpRequest(req)
	if err != nil {
		pdc_common.ReportError(err)
	}
	return req
}

func (api *SellerApi) SendRequest(req *http.Request, hasil any) error {
	var res *http.Response

	b := retry.NewFibonacci(time.Second)
	err := retry.Do(context.Background(), retry.WithMaxRetries(3, b), func(ctx context.Context) error {
		resdata, err := ClientApi.Do(req)
		res = resdata
		if err != nil {
			log.Println(api.Username, "retry", res.Request.URL.String())
			return retry.RetryableError(err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	body, _ := io.ReadAll(res.Body)
	log.Println(string(body), "asdasdasd", res.Status)
	err = json.Unmarshal(body, hasil)
	if err != nil {
		return pdc_common.ReportError(err)
	}

	return api.Session.Update(res.Cookies())
}

type YmdDate struct {
	time.Time
}

func NewEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()
	encoder.RegisterEncoder(YmdDate{}, func(v reflect.Value) string {
		value := v.Interface().(YmdDate)

		return value.Format("2006-01-02")
	})

	return encoder
}

type BasicRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
