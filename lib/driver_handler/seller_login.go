package driver_handler

import (
	"context"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func (d *DriverAccount) OpenSellerPage(ctx context.Context) error {
	asiapasifik := `(//*/div[contains(@class, "index__itemContainer")])[1]`
	button := `(//*/div[contains(@class, "index__listContainer")]/div[contains(@class, "index__siteItem")]/div[contains(@class, "index__btn")]/button)[2]`
	return chromedp.Run(ctx,
		chromedp.Navigate("https://seller.tiktok.com/"),
		chromedp.WaitReady(asiapasifik, chromedp.BySearch),
		chromedp.Click(asiapasifik, chromedp.BySearch),
		chromedp.Focus(button, chromedp.BySearch),
		chromedp.Click(button, chromedp.BySearch),
	)
}

func (d *DriverAccount) SellerLogin(dctx *DriverContext) error {
	chatsearch := `//*/div[@id="im-entry"]`

	go d.OpenSellerPage(dctx.Ctx)

	err := chromedp.Run(dctx.Ctx,
		chromedp.WaitReady(chatsearch, chromedp.BySearch),
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
			dctx.Logined = true

			return nil
		}),
	)

	return err
}
