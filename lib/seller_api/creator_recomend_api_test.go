package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/api_scenario"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
)

func TestRecommendApi(t *testing.T) {
	driver := api_scenario.UseDefaultDriver(t)
	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	payload := seller_api.NewCreatorRecomPayload()

	hasil, err := api.CreatorRecomendation(payload)
	assert.NotEmpty(t, hasil.Data)
	assert.NotEmpty(t, hasil.Data.CreatorProfileList)
	assert.Nil(t, err)

}
