package seller_api

import (
	"net/http"
)

type ProductListRes struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *ProductData `json:"data"`
}
type Image struct {
	Height       int      `json:"height"`
	Width        int      `json:"width"`
	ThumbURLList []string `json:"thumb_url_list"`
	URI          string   `json:"uri"`
	URLList      []string `json:"url_list"`
}

type SalePriceRanges struct {
	Region     string `json:"region"`
	PriceRange string `json:"price_range"`
}
type BasePrice struct {
	Region                        string `json:"region"`
	Currency                      string `json:"currency"`
	SalePrice                     string `json:"sale_price"`
	SalePriceDisplay              string `json:"sale_price_display"`
	LocalizedDutiablePrice        string `json:"localized_dutiable_price"`
	LocalizedDutiablePriceDisplay string `json:"localized_dutiable_price_display"`
}
type RegionPrices struct {
	Region                        string `json:"region"`
	Currency                      string `json:"currency"`
	SalePrice                     string `json:"sale_price"`
	SalePriceDisplay              string `json:"sale_price_display"`
	LocalizedDutiablePrice        string `json:"localized_dutiable_price"`
	LocalizedDutiablePriceDisplay string `json:"localized_dutiable_price_display"`
}
type Warehouse struct {
	WarehouseID     string `json:"warehouse_id"`
	GeoIDL0         string `json:"geo_id_l0"`
	GeoIDL1         string `json:"geo_id_l1"`
	GeoIDL2         string `json:"geo_id_l2"`
	GeoIDL3         string `json:"geo_id_l3"`
	GeoIDL4         string `json:"geo_id_l4"`
	GeoNameL0       string `json:"geo_name_l0"`
	GeoNameL1       string `json:"geo_name_l1"`
	GeoNameL2       string `json:"geo_name_l2"`
	GeoNameL3       string `json:"geo_name_l3"`
	WarehouseType   int    `json:"warehouse_type"`
	Name            string `json:"name"`
	WarehouseSource int    `json:"warehouse_source"`
}
type Quantities struct {
	WarehouseID         string    `json:"warehouse_id"`
	TotalQuantity       int       `json:"total_quantity"`
	AvailableQuantity   int       `json:"available_quantity"`
	WithholdingQuantity int       `json:"withholding_quantity"`
	Warehouse           Warehouse `json:"warehouse"`
	ReservedQuantity    int       `json:"reserved_quantity"`
	OpenQuantity        int       `json:"open_quantity"`
}
type Properties struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ValueID   string `json:"value_id"`
	ValueName string `json:"value_name"`
}
type Gtin struct {
}
type ComboSkuPositiveInfo struct {
}
type ComboSkuReverseInfo struct {
	SubSkuInComboSkuCurrent int `json:"sub_sku_in_combo_sku_current"`
}
type ComboSkuInfo struct {
	ComboSkuPositiveInfo ComboSkuPositiveInfo `json:"combo_sku_positive_info"`
	ComboSkuReverseInfo  ComboSkuReverseInfo  `json:"combo_sku_reverse_info"`
}
type Skus struct {
	ID              string         `json:"id"`
	SellerSku       string         `json:"seller_sku"`
	BasePrice       BasePrice      `json:"base_price"`
	RegionPrices    []RegionPrices `json:"region_prices"`
	Quantities      []Quantities   `json:"quantities"`
	Properties      []Properties   `json:"properties"`
	CreateTime      int            `json:"create_time"`
	PreOrderShipDay int            `json:"pre_order_ship_day"`
	Gtin            Gtin           `json:"gtin"`
	PreOrderStatus  int            `json:"pre_order_status"`
	ComboSkuInfo    ComboSkuInfo   `json:"combo_sku_info"`
}
type EditTime struct {
	UpdateTime      int `json:"update_time"`
	DeleteTime      int `json:"delete_time"`
	FreezeTime      int `json:"freeze_time"`
	CreateTime      int `json:"create_time"`
	DraftCreateTime int `json:"draft_create_time"`
}
type VisibleStatus struct {
}
type SuspendReason struct {
}
type HolidayModeStatus struct {
}
type PriceRange struct {
	MinSalePrice        string `json:"min_sale_price"`
	MaxSalePrice        string `json:"max_sale_price"`
	MinSalePriceDisplay string `json:"min_sale_price_display"`
	MaxSalePriceDisplay string `json:"max_sale_price_display"`
}
type ProductStatusView struct {
	ProductMainStatus int `json:"product_main_status"`
}
type SellerQuantity struct {
	TotalQuantity  int `json:"total_quantity"`
	OpenQuantity   int `json:"open_quantity"`
	CommitQuantity int `json:"commit_quantity"`
}
type Quantity struct {
	TotalAvailableStock int            `json:"total_available_stock"`
	SellerQuantity      SellerQuantity `json:"seller_quantity"`
}
type ComboProductPositiveInfo struct {
	IsCombo bool `json:"is_combo"`
}
type ComboProductReverseInfo struct {
	IsInCombo                       bool `json:"is_in_combo"`
	SubProductInComboProductCurrent int  `json:"sub_product_in_combo_product_current"`
}
type ComboProductInfo struct {
	ComboProductPositiveInfo ComboProductPositiveInfo `json:"combo_product_positive_info"`
	ComboProductReverseInfo  ComboProductReverseInfo  `json:"combo_product_reverse_info"`
}
type Categories struct {
	ID       string `json:"id"`
	NameKey  string `json:"name_key"`
	ParentID string `json:"parent_id"`
	Level    int    `json:"level"`
	IsLeaf   bool   `json:"is_leaf"`
	Name     string `json:"name"`
}
type Products struct {
	IsOnlineVersion     bool              `json:"is_online_version"`
	ProductID           string            `json:"product_id"`
	ProductName         string            `json:"product_name"`
	Image               Image             `json:"image"`
	SalePriceRanges     []SalePriceRanges `json:"sale_price_ranges"`
	Actions             []int             `json:"actions"`
	Skus                []Skus            `json:"skus"`
	EditTime            EditTime          `json:"edit_time"`
	VisibleStatus       VisibleStatus     `json:"visible_status"`
	TotalSkuCount       int               `json:"total_sku_count"`
	SuspendReason       SuspendReason     `json:"suspend_reason"`
	ShowcaseBindCount   int               `json:"showcase_bind_count"`
	HolidayModeStatus   HolidayModeStatus `json:"holiday_mode_status"`
	PriceRange          PriceRange        `json:"price_range"`
	ProductStatusView   ProductStatusView `json:"product_status_view"`
	TotalAvailableStock int               `json:"total_available_stock"`
	NeedLockPrice       bool              `json:"need_lock_price"`
	Quantity            Quantity          `json:"quantity"`
	ViolationRecordsID  string            `json:"violation_records_id"`
	IsMultiWarehouse    bool              `json:"is_multi_warehouse"`
	SpoBindStatus       int               `json:"spo_bind_status"`
	SellerConfirmTypes  []int             `json:"seller_confirm_types"`
	IsWithDefaultSku    bool              `json:"is_with_default_sku"`
	PreOrderStatus      int               `json:"pre_order_status"`
	ComboProductInfo    ComboProductInfo  `json:"combo_product_info"`
	Categories          []Categories      `json:"categories"`
}
type ProductData struct {
	PageNumber        int        `json:"page_number"`
	PageSize          int        `json:"page_size"`
	TotalProductCount int        `json:"total_product_count"`
	Products          []Products `json:"products"`
}
type ProductQuery struct {
	*SellerQuery
	TabID             int `schema:"tab_id"`
	PageNumber        int `schema:"page_number"`
	PageSize          int `schema:"page_size"`
	SkuNumber         int `schema:"sku_number"`
	ProductSortFields int `schema:"product_sort_fields"`
	ProductSortTypes  int `schema:"product_sort_types"`
}

func (api *SellerApi) NewProductQuery() (*ProductQuery, error) {
	selquery, err := api.NewSellerQueryWithID()
	query := ProductQuery{
		SellerQuery:       selquery,
		TabID:             2,
		PageNumber:        1,
		PageSize:          50,
		SkuNumber:         1,
		ProductSortFields: 3,
		ProductSortTypes:  0,
	}

	return &query, err

}

func (api *SellerApi) ProductList() (*ProductListRes, error) {
	var hasil ProductListRes

	ur := "https://seller-id.tiktok.com/api/v1/product/local/products/list"
	query, err := api.NewProductQuery()
	if err != nil {
		return nil, err
	}

	req := api.NewRequestJSON(http.MethodGet, ur, query, nil)
	api.SetHeader(req, map[string]string{
		"Referer": "https://seller-id.tiktok.com/product/manage?tab=active",
	})
	err = api.SendRequest(req, &hasil)

	return &hasil, err
}

type ProductCategoryRes struct {
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	CategoryInfo []*CategoryInfo `json:"category_info"`
}
type CategoryInfo struct {
	Level               int    `json:"level"`
	CategoryStarlingKey string `json:"category_starling_key"`
	CategoryID          string `json:"category_id"`
}

type CategoryProductQuery struct {
	*AffiliateQuery
	Level int `json:"level" schema:"level"`
}

func (api *SellerApi) ProductCategory() (*ProductCategoryRes, error) {
	var hasil ProductCategoryRes

	ur := "https://affiliate.tiktok.com/api/v1/affiliate/product_category/list"

	query := CategoryProductQuery{
		AffiliateQuery: api.NewAffiliateQuery(),
		Level:          1,
	}

	req := api.NewRequestJSON(http.MethodGet, ur, query, nil)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/plan/targeted/create?shop_region=ID",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ProductSearchRes struct {
	Code     int                  `json:"code"`
	Message  string               `json:"message"`
	TotalNum int                  `json:"total_num"`
	Products []*ProductSearchItem `json:"products"`
	SearchID any                  `json:"search_id"`
}

type PriceRangeSearch struct {
	MinPriceFormat string `json:"min_price_format"`
	MinPrice       string `json:"min_price"`
	MaxPriceFormat string `json:"max_price_format"`
	MaxPrice       string `json:"max_price"`
}
type PriceRangeEarn struct {
	MinPriceFormat any `json:"min_price_format"`
}
type SalePriceRange struct {
	MinPriceFormat string `json:"min_price_format"`
	MinPrice       string `json:"min_price"`
	MaxPriceFormat string `json:"max_price_format"`
	MaxPrice       string `json:"max_price"`
}
type ProductSearchItem struct {
	ProductID      string           `json:"product_id"`
	Image          Image            `json:"image"`
	Name           string           `json:"name"`
	Status         int              `json:"status"`
	PriceRange     PriceRangeSearch `json:"price_range"`
	PriceRangeEarn PriceRangeEarn   `json:"price_range_earn"`
	SalePriceRange SalePriceRange   `json:"sale_price_range"`
}

type ProductSearchPayload struct {
	SearchID        string `json:"search_id"`
	SearchKey       int    `json:"search_key"`
	KeyWord         string `json:"key_word"`
	CategoryID      string `json:"categoryId"`
	PlanType        int    `json:"plan_type"`
	PageSize        int    `json:"page_size"`
	CurPage         int    `json:"cur_page"`
	FirstCategoryID string `json:"first_category_id,omitempty"`
}

// {"search_id":"0","search_key":2,"key_word":"","categoryId":"","plan_type":2,"page_size":50,"cur_page":1,"first_category_id":"603014"}
func (api *SellerApi) ProductSearch(payload *ProductSearchPayload) (*ProductSearchRes, error) {
	var hasil ProductSearchRes

	ur := "https://affiliate.tiktok.com/api/v1/affiliate/product/search"

	query := api.NewAffiliateQuery()

	req := api.NewRequestJSON(http.MethodPost, ur, query, payload)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/plan/targeted/create?shop_region=ID",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
