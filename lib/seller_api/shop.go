package seller_api

import "net/http"

// https://seller-id.tiktok.com/api/v1/seller/common/get

type ShopCommonQuery struct {
	SellerQuery
}

type Timezone struct {
	Name string `json:"name"`
}
type Seller struct {
	Status                  int      `json:"status"`
	SellerType              int      `json:"seller_type"`
	BusinessType            int      `json:"business_type"`
	ShopName                string   `json:"shop_name"`
	SellerID                string   `json:"seller_id"`
	ShopCode                string   `json:"shop_code"`
	RegionCode              string   `json:"region_code"`
	SellerCondition         int      `json:"seller_condition"`
	ActivateCurrentTodoStep int      `json:"activate_current_todo_step"`
	IsActivateTodoFinish    bool     `json:"is_activate_todo_finish"`
	Timezone                Timezone `json:"timezone"`
	GlobalShopName          string   `json:"global_shop_name"`
	ShopRegion              string   `json:"shop_region"`
	TestFlag                int      `json:"test_flag"`
	ShopStatus              int      `json:"shop_status"`
}
type SellerMap struct {
	SellerID   string `json:"seller_id"`
	Region     string `json:"region"`
	RegionName string `json:"region_name"`
	IsNew      bool   `json:"is_new"`
}
type GlobalSeller struct {
	GlobalSellerID string      `json:"global_seller_id"`
	SellerMap      []SellerMap `json:"seller_map"`
	Status         int         `json:"status"`
	Timezone       Timezone    `json:"timezone"`
	BaseGeoID0     string      `json:"base_geo_id0"`
	SellerType     int         `json:"seller_type"`
	PartnerChannel int         `json:"partnerChannel"`
}
type DelegationMode struct {
	SellerCenterDelegationSwitchOn  bool `json:"seller_center_delegation_switch_on"`
	AffiliateDelegationModeSwitchOn bool `json:"affiliate_delegation_mode_switch_on"`
	DelegationIdentity              int  `json:"delegation_identity"`
	DelegationProtocolSigned        bool `json:"delegation_protocol_signed"`
}
type Data struct {
	Seller              Seller         `json:"seller"`
	AccountID           string         `json:"account_id"`
	IsSubAccount        bool           `json:"is_sub_account"`
	VerifiedDocuments   int            `json:"verified_documents"`
	HasOrder            bool           `json:"has_order"`
	AddedProduct        bool           `json:"added_product"`
	BindedBankAccount   bool           `json:"binded_bank_account"`
	HasUncollectedMoney bool           `json:"has_uncollected_money"`
	Conflict            bool           `json:"conflict"`
	GlobalSeller        GlobalSeller   `json:"global_seller"`
	HasStepTask         bool           `json:"has_step_task"`
	DelegationMode      DelegationMode `json:"delegation_mode"`
	UIDTestFlag         int            `json:"uid_test_flag"`
	CheckDdqPassed      bool           `json:"check_ddq_passed"`
}

type CommonInfoRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func (api *SellerApi) CommonInfo() (*CommonInfoRes, error) {
	var hasil CommonInfoRes

	ur := "https://seller-id.tiktok.com/api/v1/seller/common/get"
	query := api.NewSellerQuery()
	req := api.NewRequestJSON(http.MethodGet, ur, query, nil)
	api.SetHeader(req, map[string]string{
		"Referer": "https://seller-id.tiktok.com/product/manage",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
