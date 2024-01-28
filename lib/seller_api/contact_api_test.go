package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/api_scenario"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
)

func TestCntactApi(t *testing.T) {
	driver := api_scenario.UseDefaultDriver(t)
	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	hasil, err := api.ContactInfo(&seller_api.ContactInfoQuery{
		CreatorOecUid: "7493990942407493446",
		ShopID:        "7494567309891832821",
		Scene:         10,
	})
	t.Log(hasil)
	assert.NotEmpty(t, hasil.ContactInfo)
	assert.Nil(t, err)

}
