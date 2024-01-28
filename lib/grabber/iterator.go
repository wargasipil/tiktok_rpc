package grabber

import (
	"context"
	"time"

	"github.com/sethvargo/go-retry"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
)

type ProgressHandler interface {
	UpdateProgress(curent int, total int) error
}

type CreatorGrabber struct{}

type CreatorHandler func(page int, creator *seller_api.CreatorProfile) error

func IterateCreator(ctx context.Context, driver *driver_handler.DriverAccount, handler CreatorHandler, payload *seller_api.CreatorRecomPayload, progress ProgressHandler) error {

	api, saveSession, err := driver.CreateSellerApi()
	if err != nil {
		return err
	}
	defer saveSession()
Parent:
	for {
		var hasil *seller_api.CreatorRecomRes
		select {
		case <-ctx.Done():
			return nil
		default:
			b := retry.NewFibonacci(time.Second)
			err := retry.Do(ctx, retry.WithMaxRetries(3, b), func(ctx context.Context) error {
				data, err := api.CreatorRecomendation(payload)
				hasil = data
				if err != nil {
					return retry.RetryableError(err)
				}
				return nil
			})

			for _, creator := range hasil.Data.CreatorProfileList {
				err := handler(payload.Request.Pagination.Page, creator)
				if err != nil {
					return err
				}
			}

			nav := hasil.Data.NextPagination
			progress.UpdateProgress(payload.Request.Pagination.Page, nav.TotalPage)

			payload.Request.Pagination.Page = hasil.Data.NextPagination.NextPage

			if err != nil {
				return err
			}

			if !hasil.Data.NextPagination.HasMore {
				break Parent
			}

		}

	}

	return nil
}
