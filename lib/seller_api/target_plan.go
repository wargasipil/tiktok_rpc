package seller_api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PlanCheckCreatorRes struct {
	Response
	ConflictCreatorIds           []string          `json:"conflict_creator_ids"`
	ConflictCreatorIDToProductID map[string]string `json:"conflict_creator_id_to_product_id"`
	Conflicts                    []Conflicts       `json:"conflicts"`
}
type ImageCheck struct {
	URLList []string `json:"url_list"`
}
type ProductCheck struct {
	ProductID  string `json:"product_id"`
	Image      Image  `json:"image"`
	Name       string `json:"name"`
	Status     any    `json:"status"`
	PriceRange any    `json:"price_range"`
}
type CreatorCheck struct {
	CreatorID  string     `json:"creator_id"`
	NickName   string     `json:"nick_name"`
	Platform   string     `json:"platform"`
	Image      ImageCheck `json:"image"`
	IsVerified any        `json:"is_verified"`
	UserName   string     `json:"user_name"`
}
type Conflicts struct {
	PlanID     string       `json:"plan_id"`
	RelationID any          `json:"relation_id"`
	Product    ProductCheck `json:"product"`
	BindStatus int          `json:"bind_status"`
	Creator    CreatorCheck `json:"creator"`
	TimeZone   string       `json:"time_zone"`
	StartTime  string       `json:"start_time"`
	EndTime    string       `json:"end_time"`
}

type CreatorCheckPayload struct {
	CreatorIds     []string `json:"creator_ids"`
	ProductIds     []string `json:"product_ids"`
	PlanSourceFrom int      `json:"plan_source_from"`
}

func (api *SellerApi) PlanCheckCreator(payload *CreatorCheckPayload) (*PlanCheckCreatorRes, error) {
	var hasil PlanCheckCreatorRes

	ur := "https://affiliate.tiktok.com/api/v1/affiliate/commission_unique/check"
	query := api.NewAffiliateQuery()

	req := api.NewRequestJSON(http.MethodPost, ur, query, payload)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/plan/targeted/create?shop_region=ID",
	})
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return &hasil, err
	}
	return &hasil, nil
}

type TargetPlanCreatePayload struct {
	TargetPlans []*TargetPlans `json:"target_plans"`
}
type MetaPlans struct {
	MetaID         string `json:"meta_id"`
	MetaType       int    `json:"meta_type"`
	CommissionRate int    `json:"commission_rate"`
}
type TargetPlans struct {
	PlanName   string       `json:"plan_name"`
	EndTime    time.Time    `json:"end_time"`
	MetaPlans  []*MetaPlans `json:"meta_plans"`
	CreatorIds []string     `json:"creator_ids"`
}

func (u *TargetPlans) MarshalJSON() ([]byte, error) {
	endtime := u.EndTime.UnixMilli()
	endtimestr := strconv.FormatInt(endtime, 10)

	type Alias TargetPlans
	data, err := json.Marshal(&struct {
		EndTime string `json:"end_time"`
		*Alias
	}{
		EndTime: endtimestr,
		Alias:   (*Alias)(u),
	})
	return data, err
}

func (u *TargetPlans) UnmarshalJSON(data []byte) error {
	type Alias TargetPlans
	aux := &struct {
		EndTime string `json:"end_time"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ts, _ := strconv.ParseInt(aux.EndTime, 10, 64)
	u.EndTime = time.UnixMilli(ts)

	return nil
}

type TargetPlanCreateRes struct {
	Response
	LogicPlanIds []string `json:"logic_plan_ids"`
}

func (api *SellerApi) CreateTargetPlan(payload *TargetPlanCreatePayload) (*TargetPlanCreateRes, error) {
	var hasil TargetPlanCreateRes
	ur := "https://affiliate.tiktok.com/api/v2/affiliate/target_plan/create"
	query := api.NewAffiliateQuery()

	req := api.NewRequestJSON(http.MethodPost, ur, query, payload)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/plan/targeted/create?shop_region=ID",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

// url post
// https://affiliate.tiktok.com/api/v1/affiliate/commission_unique/check?user_language=en&shop_region=ID&aid=4331&app_name=i18n_ecom_alliance&device_id=0&fp=verify_lj1he47x_MMMISYn8_w5cP_4AtK_AYv2_IBgm43SPPHtT&device_platform=web&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=en-GB&browser_platform=Win32&browser_name=Mozilla&browser_version=5.0+(Windows+NT+10.0%3B+Win64%3B+x64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F114.0.0.0+Safari%2F537.36&browser_online=true&timezone_name=Asia%2FBangkok&msToken=YUnvaix8K9a5VuzYrHiIKiPYBMITA_lLQdQl3ap_2s8q4N7TNq9FEzFDrjfB2MX4aCKxnEmrRFhBMxG3jGhUgZV35Yva8ylWzOHEE-AMyDyQH30VRM3LyE3R0_RFcgLD&X-Bogus=DFSzswVLRAcb/2l0trFMgYT8gyYX&_signature=_02B4Z6wo00001z-f.PgAAIDCYNwTIbFYjZc.n.hAAKtz0e

// payload
// {"creator_ids":["6934006532049110018"],"product_ids":["1729640909708299253"],"plan_source_from":0}

// res
// {"code":0,"message":""}

// plan

// https://affiliate.tiktok.com/api/v2/affiliate/target_plan/create?user_language=en&shop_region=ID&aid=4331&app_name=i18n_ecom_alliance&device_id=0&fp=verify_lj1he47x_MMMISYn8_w5cP_4AtK_AYv2_IBgm43SPPHtT&device_platform=web&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=en-GB&browser_platform=Win32&browser_name=Mozilla&browser_version=5.0+(Windows+NT+10.0%3B+Win64%3B+x64)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F114.0.0.0+Safari%2F537.36&browser_online=true&timezone_name=Asia%2FBangkok&msToken=1ec9jDK_-D7zR4wa5tiZWJOgdcYRIQPuwsk-7uM3h0_J03p4CNZ7AMnXoWmkhsHpCH2MCmGEcSRWLwTt1XskPtW-Hv8K4e-oY04KLSgToqs4ru49A5mJ2X9Ks4QhBz05&X-Bogus=DFSzswVLtgrXeGl0trFKQKT8gyOT&_signature=_02B4Z6wo00001NOBFcwAAIDBjML6FYMhwVjTgRFAAFB001

// pay
// {"target_plans":[{"plan_name":"plan test","end_time":"1688144399000","meta_plans":[{"meta_id":"1729640909708299253","meta_type":1,"commission_rate":1000}],"creator_ids":["6934006532049110018"]}]}

// res

// {"code":0,"message":"","logic_plan_ids":["9527432005"]}

type ErrorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *ErrorRes) Error() string {
	return err.Message
}

type CheckUsernameRes struct {
	ErrorRes
	TotalNum int              `json:"total_num"`
	Creators []*CheckCreators `json:"creators"`
	SearchID any              `json:"search_id"`
}
type CheckImage struct {
	URI     string   `json:"uri"`
	URLList []string `json:"url_list"`
}
type CheckCreators struct {
	CreatorID            string      `json:"creator_id"`
	CreatorBio           string      `json:"creator_bio"`
	NickName             string      `json:"nick_name"`
	Platform             string      `json:"platform"`
	Image                *CheckImage `json:"image"`
	IsVerified           bool        `json:"is_verified"`
	UserName             string      `json:"user_name"`
	CreatorOecID         string      `json:"creator_oec_id"`
	AffiliateCreatorType int         `json:"affiliate_creator_type"`
}

type PlanCheckUserPayload struct {
	PageSize  int    `json:"page_size"`
	SearchID  string `json:"search_id"`
	SearchKey int    `json:"search_key"`
	KeyWord   string `json:"key_word"`
}

func (api *SellerApi) PlanCheckUsername(username string) (*CheckUsernameRes, error) {
	var hasil CheckUsernameRes

	payload := PlanCheckUserPayload{
		PageSize:  20,
		SearchID:  "0",
		SearchKey: 4,
		KeyWord:   username,
	}

	ur := "https://affiliate.tiktok.com/api/v1/affiliate/creator/search"
	query := api.NewAffiliateQuery()

	req := api.NewRequestJSON(http.MethodPost, ur, query, payload)
	api.SetHeader(req, map[string]string{
		"Referer": "https://affiliate.tiktok.com/plan/targeted/create?shop_region=ID",
	})
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

// {"page_size":20,"search_id":"0","search_key":4,"key_word":"dwiipa6"}
