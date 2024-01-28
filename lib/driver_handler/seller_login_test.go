package driver_handler_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
)

func TestLogin(t *testing.T) {
	driver, err := driver_handler.NewDriverAccount("test@gmail.com", "", "", "")
	assert.Nil(t, err)

	driver.Run(false, func(dctx *driver_handler.DriverContext) error {
		driver.SellerLogin(dctx)

		return nil
	})
}
