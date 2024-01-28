package seller_api

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type BasicQuery struct {
	AID             int    `schema:"aid"`
	AppName         string `schema:"app_name"`
	DeviceID        int    `schema:"device_id"`
	Fp              string `schema:"fp"`
	DevicePlatform  string `schema:"device_platform"`
	CookieEnabled   bool   `schema:"cookie_enabled"`
	ScreenWidth     int    `schema:"screen_width"`
	ScreenHeight    int    `schema:"screen_height"`
	BrowserLanguage string `schema:"browser_language"`
	BrowserPlatform string `schema:"browser_platform"`
	BrowserName     string `schema:"browser_name"`
	BrowserVersion  string `schema:"browser_version"`
	BrowserOnline   bool   `schema:"browser_online"`
	TimezoneName    string `schema:"timezone_name"`
	MsToken         string `schema:"msToken"`
	XBogus          string `schema:"X-Bogus"`
	Signature       string `schema:"_signature"`
}

func (api *SellerApi) NewBasicQuery(appid int, appName string) BasicQuery {
	fp := genVerifyFp()
	token := api.Session.FindCookie("msToken")
	ua := api.Session.UserAgent()
	uas := strings.SplitN(ua, `/`, 2)

	basicQuery := BasicQuery{
		AID:             appid,
		AppName:         appName,
		DeviceID:        0,
		Fp:              fp,
		DevicePlatform:  "web",
		CookieEnabled:   true,
		ScreenWidth:     1920,
		ScreenHeight:    1080,
		BrowserLanguage: "en-GB",
		BrowserPlatform: "Win32",
		BrowserName:     uas[0],
		BrowserVersion:  uas[1],
		BrowserOnline:   true,
		TimezoneName:    "Asia/Bangkok",
		MsToken:         token,
		XBogus:          "",
		Signature:       "",
	}
	return basicQuery
}

type AffiliateQuery struct {
	BasicQuery
	UserLanguage string `schema:"user_language"`

	ShopRegion string `schema:"shop_region"`
}

func (api *SellerApi) NewAffiliateQuery() *AffiliateQuery {
	base := api.NewBasicQuery(4331, "i18n_ecom_alliance")
	query := AffiliateQuery{
		BasicQuery:   base,
		UserLanguage: "en",

		ShopRegion: "ID",
	}

	return &query
}

type SellerQuery struct {
	BasicQuery
	Locale      string `schema:"locale"`
	Language    string `schema:"language"`
	OecSellerID string `schema:"oec_seller_id"`
}

func (api *SellerApi) NewSellerQuery() *SellerQuery {
	base := api.NewBasicQuery(4068, "i18n_ecom_shop")
	query := SellerQuery{
		BasicQuery: base,
		Locale:     "en",
		Language:   "en",
	}

	return &query
}

func (api *SellerApi) NewSellerQueryWithID() (*SellerQuery, error) {
	query := api.NewSellerQuery()
	seller, err := api.Data.Get()
	if err != nil {
		return nil, err
	}
	query.OecSellerID = seller.OecSellerID
	return query, nil
}

// query shop

// locale: en
// language: en
// oec_seller_id: 7494567309891832821
// aid: 4068
// app_name: i18n_ecom_shop
// device_id: 0
// fp: verify_lj1b3j3m_ayO3nQ8Z_BDvU_4XIf_ATj8_soAy5Dwkp2J5
// device_platform: web
// cookie_enabled: true
// screen_width: 1920
// screen_height: 1080
// browser_language: en-GB
// browser_platform: Win32
// browser_name: Mozilla
// browser_version: 5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36
// browser_online: true
// timezone_name: Asia/Bangkok
// tab_id: 2
// page_number: 1
// page_size: 50
// sku_number: 1
// product_sort_fields: 3
// product_sort_types: 0
// msToken: rcMZWVG9dcmSnze2p-kDQ-Vmq3BYWBWj0pm4BhpYOqBsdARvuGFGxIQy7nYXMrnyeC00Zy9kfCCUvD6kG-JYa0OyEt8sp0f9yULWox0aanDWjV7ydjzqGQ==
// X-Bogus: DFSzswVL/8xANGl0trYqeYT8gyYq
// _signature: _02B4Z6wo00001kV10UwAAIDDGjY-lZsTpZZFddXAAPXi58

// need_verify_account: true

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func base36Encode(number int64) string {
	alphabets := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base36 := ""
	sign := ""

	if number < 0 {
		sign = "-"
		number = -number
	}
	if 0 <= number && int(number) < len(alphabets) {
		return sign + string(alphabets[number])
	}

	for number != 0 {
		newNumber, i := divmod(number, int64(len(alphabets)))
		base36 = string(alphabets[i]) + base36
		number = newNumber
	}

	return sign + base36
}

func genVerifyFp() string {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"[:]
	chars_len := len(chars)

	now := time.Now()
	scenario_title := base36Encode(now.UnixMilli())
	var uuid [36]string
	uuid[8] = "_"
	uuid[13] = "_"
	uuid[18] = "_"
	uuid[23] = "_"
	uuid[14] = "4"

	for i := 1; i < 36; i++ {
		if uuid[i] != "" {
			continue
		}
		r := int(rand.Float64() * float64(chars_len))
		var a int
		if i == 19 {
			a = 8
		} else {
			a = r
		}
		uuid[i] = string(chars[int((3&r)|a)])
	}
	return fmt.Sprintf("verify_%s_%s", strings.ToLower(scenario_title), strings.Join(uuid[:], ""))
}
