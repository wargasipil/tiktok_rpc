package seller_api_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/api_scenario"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
)

var ProductID = "1729720596544392181"

func TestCheckCreator(t *testing.T) {
	driver, err := driver_handler.NewDriverAccount("test@gmail.com", "", "", "")
	assert.Nil(t, err)
	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	payload := seller_api.CreatorCheckPayload{
		CreatorIds:     []string{"6800259471647196162"},
		ProductIds:     []string{"1729640909708299253"},
		PlanSourceFrom: 0,
	}

	hasil, err := api.PlanCheckCreator(&payload)
	assert.Equal(t, hasil.Code, 0)
	assert.Nil(t, err)
}

func TestCreatePlan(t *testing.T) {
	driver := api_scenario.UseDefaultDriver(t)

	// driver.Run(false, func(dctx *driver_handler.DriverContext) error {
	// 	driver.SellerLogin(dctx)

	// 	return chromedp.Run(dctx.Ctx,
	// 		chromedp.Navigate("https://affiliate.tiktok.com/plan/targeted"),
	// 	)
	// })
	// assert.Nil(t, err)

	api, saveSession, err := driver.CreateSellerApi()
	assert.Nil(t, err)
	defer saveSession()

	// datastr := `{"target_plans":[{"end_time":"1689613200000","plan_name":"defaultas3","meta_plans":[{"meta_id":"1729640909708299253","meta_type":1,"commission_rate":1000}],"creator_ids":["7493998798323746316"]}]}`

	datastr := `{"target_plans":[{"plan_name":"asdasdwaasw","end_time":"1697529293000","meta_plans":[{"meta_id":"1729640909708299253","meta_type":1,"commission_rate":1200}],"creator_ids":["6607561898710532097"]}]}`

	var payload seller_api.TargetPlanCreatePayload
	json.Unmarshal([]byte(datastr), &payload)

	payload.TargetPlans[0].MetaPlans[0].MetaID = ProductID

	// {"target_plans":[{"plan_name":"plan test","end_time":"1688144399000","meta_plans":[{"meta_id":"1729640909708299253","meta_type":1,"commission_rate":1000}],"creator_ids":["6934006532049110018"]}]}

	hasil, err := api.CreateTargetPlan(&payload)
	t.Log(hasil)
	assert.Equal(t, hasil.Code, 0)
	assert.Nil(t, err)
}

// kesalahan produk 98001004
// kreator ada yang sudah ditambahkan 16003003
