package repo

// creator_id	nickname	name	follower	average_view	videos	follower_top_age_group	follower_top_gender	follower_top_gender_share	follower_top_age_group_share	video_pub_cnt	product_cnt	category	region	email	whatsapp

type CreatorTaskStatus string

const (
	CTaskRunning   CreatorTaskStatus = "running"
	CTaskCompleted CreatorTaskStatus = "completed"
	CTaskStopped   CreatorTaskStatus = "stopped"
)

type Collection struct {
	CollName   string            `json:"collection_name" gorm:"primaryKey;autoIncrement:false"`
	Status     CreatorTaskStatus `json:"status"`
	Count      int               `json:"count"`
	LimitCount int               `json:"limit_count"`
	ErrMessage string            `json:"err_message"`
	Progress   float32           `json:"progress"`
}

func NewCollection() *Collection {
	return &Collection{
		Status: CTaskRunning,
	}
}

type CreatorItem struct {
	CreatorID                string  `json:"creator_id" gorm:"primaryKey;autoIncrement:false"`
	Nickname                 string  `json:"nickname"`
	Name                     string  `json:"name"`
	Follower                 int     `json:"follower"`
	AverageView              int     `json:"average_view"`
	Videos                   int     `json:"videos"`
	FollowerTopAgeGroup      int     `json:"follower_top_age_group"`
	FollowerTopGender        int     `json:"follower_top_gender"`
	FollowerTopGenderShare   float32 `json:"follower_top_gender_share"`
	FollowerTopAgeGroupShare float32 `json:"follower_top_age_group_share"`
	VideoPubCnt              int     `json:"video_pub_cnt"`
	ProductCnt               int     `json:"product_cnt"`
	// Category                 []string
	Region     string        `json:"region"`
	Email      string        `json:"email"`
	Whatsapp   string        `json:"whatsapp"`
	Collection []*Collection `gorm:"many2many:collection;"`
}
