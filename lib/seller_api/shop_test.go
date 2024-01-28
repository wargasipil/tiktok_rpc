package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/api_scenario"
)

func TestCommonGet(t *testing.T) {
	driver := api_scenario.UseDefaultDriver(t)
	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	hasil, err := api.CommonInfo()
	t.Log(hasil)
	assert.NotEmpty(t, hasil.Data)
	assert.Nil(t, err)

}
