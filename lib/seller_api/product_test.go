package seller_api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
)

func TestProductApi(t *testing.T) {
	driver, err := driver_handler.NewDriverAccount("test@gmail.com", "", "", "")
	assert.Nil(t, err)

	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	t.Run("test product list", func(t *testing.T) {
		hasil, err := api.ProductList()
		t.Log(hasil)
		assert.NotEmpty(t, hasil.Data)
		assert.Nil(t, err)
	})

	t.Run("test product category", func(t *testing.T) {
		hasil, err := api.ProductCategory()
		t.Log(hasil)
		assert.NotEmpty(t, hasil.CategoryInfo)
		assert.Nil(t, err)
	})

	t.Run("test search product", func(t *testing.T) {
		// {"search_id":"0","search_key":2,"key_word":"","categoryId":"","plan_type":2,"page_size":50,"cur_page":1,"first_category_id":"603014"}
		payload := seller_api.ProductSearchPayload{
			SearchID:   "0",
			SearchKey:  2,
			KeyWord:    "",
			CategoryID: "",
			PlanType:   2,
			PageSize:   50,
			CurPage:    1,
		}

		hasil, err := api.ProductSearch(&payload)
		t.Log(hasil)
		assert.NotEmpty(t, hasil.Products)
		assert.Nil(t, err)
	})

}
