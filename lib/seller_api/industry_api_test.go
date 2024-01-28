package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/api_scenario"
)

func TestIndustryApi(t *testing.T) {
	driver := api_scenario.UseDefaultDriver(t)
	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	hasil, err := api.GetIndustries()
	t.Log(hasil)
	assert.NotEmpty(t, hasil.MainIndustries)
	assert.Nil(t, err)

}
