package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
)

func TestQuery(t *testing.T) {
	driver, err := driver_handler.NewDriverAccount("asasdasd", "", "", "")
	assert.Nil(t, err)
	api, saveSession, err := driver.CreateSellerApi()
	defer saveSession()

	assert.Nil(t, err)

	query := api.NewAffiliateQuery()
	assert.NotEmpty(t, query.MsToken)

}
