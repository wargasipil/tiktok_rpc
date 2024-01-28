package driver_handler

import (
	"context"
	"net/http"

	"github.com/chromedp/cdproto/network"
)

type DriverSession interface {
	SetCookieToDriver(ctx context.Context) error
	Load() error
	DeleteSession() error
	SaveSession() error
	Sync() error
	Update(cookies []*http.Cookie) error
	AddToHttpRequest(req *http.Request)
	UserAgent() string
	SaveFromDriver(cookies []*network.Cookie, ua string) error
	GetCookies() []*http.Cookie
	FindCookie(string) string
}
