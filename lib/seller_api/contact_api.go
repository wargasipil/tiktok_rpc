package seller_api

import (
	"errors"
	"net/http"
)

type ContactInfoQuery struct {
	*AffiliateQuery
	CreatorOecUid string `schema:"creator_oecuid"`
	ShopID        string `schema:"shop_id"`
	Scene         int    `schema:"scene"`
}

type ContactInfo struct {
	Field int    `json:"field"`
	Value string `json:"value"`
	Title any    `json:"title"`
}
type ContactInfoRes struct {
	*BasicRes
	ContactInfo []*ContactInfo `json:"contact_info"`
}

var ErrPhoneNotFound = errors.New("contact not found")

func (contact *ContactInfoRes) GetPhone() (string, error) {

	for _, cntc := range contact.ContactInfo {
		if cntc.Field == 1 {
			return cntc.Value, nil
		}
	}

	return "", ErrPhoneNotFound
}

// creator_oecuid: 7493990942407493446
// shop_id: 7494567309891832821
// shop_region: ID
// scene: 10

func (api *SellerApi) ContactInfo(query *ContactInfoQuery) (*ContactInfoRes, error) {
	var hasil ContactInfoRes

	query.AffiliateQuery = api.NewAffiliateQuery()
	query.Scene = 10

	ur := "https://affiliate.tiktok.com/api_sens/v1/affiliate/cmp/contact"
	req := api.NewRequestJSON(http.MethodGet, ur, query, nil)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/connection/creator",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
