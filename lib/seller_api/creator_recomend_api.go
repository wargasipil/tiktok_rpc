package seller_api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CreatorRecomPage struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type FollowerCountFilter struct {
	FollowerCntMax int `json:"follower_cnt_max,omitempty"`
	FollowerCntMin int `json:"follower_cnt_min,omitempty"`
}
type CreatorRecomReq struct {
	*FollowerCountFilter
	FollowerGenders          []any            `json:"follower_genders"`
	FollowerAgeGroups        []any            `json:"follower_age_groups"`
	ManagedByAgency          []any            `json:"managed_by_agency"`
	Pagination               CreatorRecomPage `json:"pagination"`
	CreatorScoreRange        []int            `json:"creator_score_range"`
	ContentPreferenceRange   []int            `json:"content_preference_range"`
	IsActiveCreator          bool             `json:"is_active_creator,omitempty"`
	IsHighSampleDispatchRate bool             `json:"is_high_sample_dispatch_rate,omitempty"`
	IsQuicklyResponse        bool             `json:"is_quickly_response,omitempty"`
	CreatorProductCategories []string         `json:"creator_product_categories,omitempty"`
}

type CreatorRecomPayload struct {
	Request CreatorRecomReq `json:"request"`
}

func NewCreatorRecomPayload() *CreatorRecomPayload {
	return &CreatorRecomPayload{
		Request: CreatorRecomReq{
			FollowerGenders:        []any{},
			FollowerAgeGroups:      []any{},
			ManagedByAgency:        []any{},
			Pagination:             CreatorRecomPage{Page: 0, Size: 20},
			CreatorScoreRange:      []int{4},
			ContentPreferenceRange: []int{1},
		},
	}
}

type NextPagination struct {
	HasMore        bool   `json:"has_more"`
	NextPage       int    `json:"next_page"`
	TotalPage      int    `json:"total_page"`
	Total          int    `json:"total"`
	SearchKey      string `json:"search_key"`
	NextItemCursor int    `json:"next_item_cursor"`
}
type CreatorOecuid struct {
	Value        string `json:"value"`
	IsAuthorized bool   `json:"is_authorized"`
	Status       int    `json:"status"`
}
type Handle struct {
	Value        string `json:"value"`
	IsAuthorized bool   `json:"is_authorized"`
	Status       int    `json:"status"`
}
type Nickname struct {
	Value        string `json:"value"`
	IsAuthorized bool   `json:"is_authorized"`
	Status       int    `json:"status"`
}

type Avatar struct {
	Value        AvatarValue `json:"value"`
	IsAuthorized bool        `json:"is_authorized"`
	Status       int         `json:"status"`
}
type SelectionRegion struct {
	Value        string `json:"value"`
	IsAuthorized bool   `json:"is_authorized"`
	Status       int    `json:"status"`
}
type IsOfficialRecommend struct {
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type IsShowRecomIcon struct {
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type IsOpenAccount struct {
	Value        bool `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type ProductCategories struct {
	ID          any      `json:"id"`
	CategoryIds []string `json:"category_ids"`
}
type MainIndustryValue struct {
	StarlingKey       string              `json:"starling_key"`
	ProductCategories []ProductCategories `json:"product_categories"`
	Name              string              `json:"name"`
}
type MainIndustry struct {
	Value        []MainIndustryValue `json:"value"`
	IsAuthorized bool                `json:"is_authorized"`
	Status       int                 `json:"status"`
}
type EcLiveAvgUv struct {
	Value        string `json:"value"`
	IsAuthorized bool   `json:"is_authorized"`
	Status       int    `json:"status"`
}
type IsQuicklyResponse struct {
	Value        bool `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type IsActiveCreator struct {
	Value        bool `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type IsHighSampleDispatchRate struct {
	Value        bool `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type CreatorEcLevelScore struct {
	Value        float64 `json:"value"`
	IsAuthorized bool    `json:"is_authorized"`
	Status       int     `json:"status"`
}
type FollowerCnt struct {
	Value        int  `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}

func (u *FollowerCnt) UnmarshalJSON(data []byte) error {
	type Alias FollowerCnt
	aux := &struct {
		Value string `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Value != "" {
		ts, err := strconv.Atoi(aux.Value)
		if err != nil {
			return err
		}
		u.Value = ts
	}

	return nil
}

type EcVideoGpmReference struct {
	Value        bool `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
}
type EcLiveGpmReference struct {
	Value        bool `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
}
type CreatorBindMcnName struct {
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type Value struct {
	Minimal       string `json:"minimal"`
	Maximum       string `json:"maximum"`
	Symbol        string `json:"symbol"`
	MinimalFormat string `json:"minimal_format"`
	MaximumFormat string `json:"maximum_format"`
}
type EcLiveGpm struct {
	Value        Value `json:"value"`
	IsAuthorized bool  `json:"is_authorized"`
	Status       int   `json:"status"`
}
type EcVideoAvgViewCnt struct {
	Value        int  `json:"value"`
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}

func (u *EcVideoAvgViewCnt) UnmarshalJSON(data []byte) error {
	type Alias EcVideoAvgViewCnt
	aux := &struct {
		Value string `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Value != "" {
		ts, err := strconv.Atoi(aux.Value)
		if err != nil {
			return err
		}
		u.Value = ts
	}

	return nil
}

type EcVideoGpm struct {
	Value        Value `json:"value"`
	IsAuthorized bool  `json:"is_authorized"`
	Status       int   `json:"status"`
}
type OccurredTopRank struct {
	IsAuthorized bool `json:"is_authorized"`
	Status       int  `json:"status"`
}
type Nickname0 struct {
	Value        string `json:"value"`
	IsAuthoRized bool   `json:"is_authorized"`
	Status       int    `json:"status"`
}
type AvatarValue struct {
	ThumbURLList []string `json:"thumb_url_list"`
	URLList      []string `json:"url_list"`
}

type CreatorProfile struct {
	CreatorOecuid            CreatorOecuid            `json:"creator_oecuid"`
	Handle                   Handle                   `json:"handle"`
	Nickname                 Nickname                 `json:"nickname,omitempty"`
	Avatar                   Avatar                   `json:"avatar,omitempty"`
	SelectionRegion          SelectionRegion          `json:"selection_region"`
	IsOfficialRecommend      IsOfficialRecommend      `json:"is_official_recommend"`
	IsShowRecomIcon          IsShowRecomIcon          `json:"is_show_recom_icon"`
	IsOpenAccount            IsOpenAccount            `json:"is_open_account"`
	MainIndustry             MainIndustry             `json:"main_industry"`
	EcLiveAvgUv              EcLiveAvgUv              `json:"ec_live_avg_uv"`
	IsQuicklyResponse        IsQuicklyResponse        `json:"is_quickly_response"`
	IsActiveCreator          IsActiveCreator          `json:"is_active_creator"`
	IsHighSampleDispatchRate IsHighSampleDispatchRate `json:"is_high_sample_dispatch_rate"`
	CreatorEcLevelScore      CreatorEcLevelScore      `json:"creator_ec_level_score"`
	FollowerCnt              FollowerCnt              `json:"follower_cnt"`
	EcVideoGpmReference      EcVideoGpmReference      `json:"ec_video_gpm_reference"`
	EcLiveGpmReference       EcLiveGpmReference       `json:"ec_live_gpm_reference"`
	CreatorBindMcnName       CreatorBindMcnName       `json:"creator_bind_mcn_name"`
	EcLiveGpm                EcLiveGpm                `json:"ec_live_gpm"`
	EcVideoAvgViewCnt        EcVideoAvgViewCnt        `json:"ec_video_avg_view_cnt"`
	EcVideoGpm               EcVideoGpm               `json:"ec_video_gpm"`
	OccurredTopRank          OccurredTopRank          `json:"occurred_top_rank"`
}
type CreatorRecomData struct {
	NextPagination     NextPagination    `json:"next_pagination"`
	CreatorProfileList []*CreatorProfile `json:"creator_profile_list"`
}

type CreatorRecomRes struct {
	*BasicRes
	Data *CreatorRecomData `json:"data"`
}

func (api *SellerApi) CreatorRecomendation(payload *CreatorRecomPayload) (*CreatorRecomRes, error) {
	var hasil CreatorRecomRes

	ur := "https://affiliate.tiktok.com/api/v1/oec/affiliate/creator/marketplace/recommendation"
	query := api.NewAffiliateQuery()
	req := api.NewRequestJSON(http.MethodPost, ur, query, payload)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/connection/creator",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
