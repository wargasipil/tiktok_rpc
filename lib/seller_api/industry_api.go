package seller_api

import "net/http"

type MainIndustriesPayload struct{}

type ProductCategory struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	StarlingKey string   `json:"starling_key"`
	CategoryIds []string `json:"category_ids"`
}

type MainIndustries struct {
	StarlingKey       string             `json:"starling_key"`
	ProductCategories []*ProductCategory `json:"product_categories"`
	Name              string             `json:"name"`
}

type MainIndustriesRes struct {
	Code           int               `json:"code"`
	Message        string            `json:"message"`
	MainIndustries []*MainIndustries `json:"mainIndustries"`
}

func (api *SellerApi) GetIndustries() (*MainIndustriesRes, error) {
	var hasil MainIndustriesRes
	payload := MainIndustriesPayload{}

	ur := "https://affiliate.tiktok.com/api/v1/oec/affiliate/cmp/main/industries"
	query := api.NewAffiliateQuery()
	req := api.NewRequestJSON(http.MethodPost, ur, query, payload)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/connection/creator",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
