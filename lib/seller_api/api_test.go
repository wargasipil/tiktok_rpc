package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
)

func TestQueryWithID(t *testing.T) {
	driver, err := driver_handler.NewDriverAccount("asasdasd", "", "", "")
	assert.Nil(t, err)
	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	hasil, err := api.NewSellerQueryWithID()
	t.Log(hasil)
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)

}
