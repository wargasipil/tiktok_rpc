package driver_handler

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/wargasipil/tiktok_rpc/lib/session_handler"
)

type DriverContext struct {
	sync.Mutex
	Logined bool
	Ctx     context.Context
}

type DriverAccount struct {
	Username      string
	Password      string
	Email         string
	EmailPassword string
	Headless      bool
	Proxy         string
	DevMode       bool
	ParentContext context.Context
	Session       DriverSession
	FirstProfile  bool
}

type BrowserClosed struct {
	sync.Mutex
	Data bool
}

func (d *DriverAccount) ProfilePath() string {
	pathdata, _ := filepath.Abs("/tiktok_profile/" + d.Username)
	return pathdata
}

func (d *DriverAccount) CreateContext(headless bool) (*DriverContext, func()) {
	pathProfile := d.ProfilePath()
	if _, err := os.Stat(pathProfile); errors.Is(err, os.ErrNotExist) {
		d.FirstProfile = true
		err := os.MkdirAll(pathProfile, os.ModeDir)
		if err != nil {
			panic(err)
		}
		firstfile := filepath.Join(pathProfile, "First Run")
		file, _ := os.OpenFile(firstfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		file.Close()
	}
	opt := []func(*chromedp.ExecAllocator){
		chromedp.Flag("headless", headless),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"),
		// chromedp.UserDataDir(pathProfile),
		// chromedp.Flag("profile-directory", "Default"),
	}

	if d.DevMode {
		opt = append(opt, chromedp.Flag("auto-open-devtools-for-tabs", true))
	}

	if d.Proxy != "" {
		opt = append(opt, chromedp.ProxyServer(d.Proxy))
	}

	// jika pengen custom context dari mana
	var parentCtx context.Context
	if d.ParentContext == nil {
		parentCtx = context.Background()
	} else {
		parentCtx = d.ParentContext
	}

	ctxall, cancelAloc := chromedp.NewExecAllocator(
		parentCtx,
		opt...,
	)

	ctx, cancelCtx := chromedp.NewContext(ctxall)

	dctx := DriverContext{
		Logined: false,
		Ctx:     ctx,
	}

	d.Session.SetCookieToDriver(dctx.Ctx)

	// checking jaga2 jika close manual browser nya
	isClosed := BrowserClosed{
		Data: false,
	}
	go func() {
		<-ctx.Done()

		isClosed.Lock()
		defer isClosed.Unlock()

		isClosed.Data = true
	}()

	return &dctx, func() {
		isClosed.Lock()
		defer isClosed.Unlock()

		if isClosed.Data {
			return
		}
		d.SaveSession(&dctx)
		cancelCtx()
		cancelAloc()

	}
}

func (d *DriverAccount) SaveSession(dctx *DriverContext) error {
	return chromedp.Run(dctx.Ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			var userAgent string
			err = chromedp.Evaluate("navigator.userAgent", &userAgent).Do(ctx)
			if err != nil {
				return err
			}

			err = d.Session.SaveFromDriver(cookies, userAgent)
			if err != nil {
				return err
			}
			return nil
		}),
	)

}

func (d *DriverAccount) Run(headless bool, actionCallback func(dctx *DriverContext) error) error {
	dctx, cancel := d.CreateContext(headless)
	defer cancel()

	return actionCallback(dctx)

}

func (d *DriverAccount) RunWithTimeout(headless bool, actionCallback func(keep func(), dctx *DriverContext) error, duration time.Duration) error {
	dctx, cancelCtx := d.CreateContext(headless)
	defer cancelCtx()

	activity := make(chan int)

	go func() {
		ticker := time.NewTimer(duration)

	Parent:
		for {
			select {
			case <-dctx.Ctx.Done():
				break Parent
			case <-activity:
				if !ticker.Stop() {
					<-ticker.C
				}
				ticker.Reset(duration)
			case <-ticker.C:
				log.Println("closing driver for no activity")
				cancelCtx()
			}
		}
	}()

	keep := func() {
		activity <- 1
	}
	return actionCallback(keep, dctx)
}

func RunInInconigto(handler func(ctx context.Context)) {

	ctxall, cancelAloc := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("incognito", true),
	)

	ctx, cancelCtx := chromedp.NewContext(ctxall)
	defer cancelAloc()
	defer cancelCtx()

	handler(ctx)
}

func NewDriverAccount(username string, password string, email string, pwdemail string) (*DriverAccount, error) {
	sess := session_handler.NewSession(username)
	err := sess.Load()

	if errors.Is(err, session_handler.ErrSessionNotFound) {
		err = nil
	}

	return &DriverAccount{
		Session:       sess,
		Username:      username,
		Password:      password,
		Email:         email,
		EmailPassword: pwdemail,
		Headless:      true,
		Proxy:         "",
	}, nil

}
