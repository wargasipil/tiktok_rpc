package session_handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/common_conf/pdc_common"
)

var BaseSessionPath = "/tiktok_data_session/"

func init() {
	pathdata, _ := filepath.Abs(BaseSessionPath)
	if _, err := os.Stat(pathdata); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(pathdata, os.ModeDir)
	}
}

var ErrSessionNotFound = errors.New("session tidak ada")

type Session struct {
	sync.Mutex
	fname   string
	Cookies []*http.Cookie
	Ua      string
}

func (sess *Session) UserAgent() string {
	return sess.Ua
}

func (sess *Session) GetCookies() []*http.Cookie {
	return sess.Cookies
}

func (sess *Session) Update(cookies []*http.Cookie) error {
	sess.Lock()
	defer sess.Unlock()

	for _, cookie := range cookies {
		err := sess.updateCookie(cookie)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sess *Session) updateCookie(cookie *http.Cookie) error {

	fixCookies := []*http.Cookie{}

	for _, oldCookie := range sess.Cookies {
		if oldCookie.Name == cookie.Name {
			fixCookies = append(fixCookies, cookie)
		} else {
			fixCookies = append(fixCookies, oldCookie)
		}
	}
	sess.Cookies = fixCookies
	return nil
}

func (sess *Session) SetCookieToDriver(ctx context.Context) error {
	sess.Lock()
	defer sess.Unlock()

	return chromedp.Run(
		ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {

			for _, cookie := range sess.Cookies {
				expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))

				err := network.SetCookie(cookie.Name, cookie.Value).
					WithDomain(cookie.Domain).
					WithPath(cookie.Path).
					WithHTTPOnly(cookie.HttpOnly).
					WithSecure(cookie.Secure).
					WithExpires(&expr).
					Do(ctx)

				if err != nil {
					if !errors.Is(context.Canceled, err) {
						pdc_common.ReportError(err)
					}

				}
			}
			return nil
		}),
	)
}

func (sess *Session) AddToHttpRequest(req *http.Request) {
	for _, cookie := range sess.Cookies {
		req.AddCookie(cookie)
	}
}

func (sess *Session) SaveFromDriver(cookies []*network.Cookie, ua string) error {
	sess.Lock()
	defer sess.Unlock()

	fixCookies := make([]*http.Cookie, len(cookies))

	for ind, ncookie := range cookies {
		fixCookies[ind] = &http.Cookie{
			Name:     ncookie.Name,
			Value:    ncookie.Value,
			Path:     ncookie.Path,
			Domain:   ncookie.Domain,
			Expires:  time.Unix(int64(ncookie.Expires), 0),
			Secure:   ncookie.Secure,
			HttpOnly: ncookie.HTTPOnly,
		}
	}

	sess.Cookies = fixCookies
	sess.Ua = ua
	return sess.save()

}

func (sess *Session) Sync() error {
	return nil
}

func (sess *Session) SaveSession() error {
	sess.Lock()
	defer sess.Unlock()

	return sess.save()
}

func (sess *Session) FindCookie(key string) string {
	for _, cookie := range sess.Cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}
	return ""
}

func (sess *Session) save() error {
	pathdata := filepath.Join(BaseSessionPath, sess.fname+".json")
	file, err := os.OpenFile(pathdata, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return pdc_common.ReportError(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(sess)
	if err != nil {
		return pdc_common.ReportError(err)
	}

	return nil
}

func (sess *Session) Load() error {
	sess.Lock()
	defer sess.Unlock()

	pathdata := filepath.Join(BaseSessionPath, sess.fname+".json")

	if _, err := os.Stat(pathdata); errors.Is(err, os.ErrNotExist) {
		return ErrSessionNotFound
	}

	file, err := os.ReadFile(pathdata)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, sess)
	if err != nil {
		return err
	}

	return nil
}

func (sess *Session) DeleteSession() error {
	pathdata := filepath.Join(BaseSessionPath, sess.fname+".json")
	return os.Remove(pathdata)

}

func NewSession(fname string) *Session {
	session := Session{
		fname:   fname,
		Cookies: []*http.Cookie{},
		Ua:      "",
	}
	return &session
}
