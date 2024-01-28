package driver_handler

import (
	"errors"

	"github.com/wargasipil/tiktok_rpc/lib/repo"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
	"github.com/wargasipil/tiktok_rpc/lib/session_handler"
)

func (d *DriverAccount) CreateSellerApi() (*seller_api.SellerApi, func() error, error) {
	err := d.Session.Load()

	if errors.Is(err, session_handler.ErrSessionNotFound) {
		d.Run(false, func(dctx *DriverContext) error {
			return d.SellerLogin(dctx)
		})
	}

	cache := &repo.AccountCache{}
	return seller_api.NewSellerApi(d.Username, d.Session, cache)

}
