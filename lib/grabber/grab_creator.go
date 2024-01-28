package grabber

import (
	"context"
	"errors"
	"log"

	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
	"github.com/wargasipil/tiktok_rpc/lib/repo"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
	"gorm.io/gorm"
)

type GrabCreatorConfig struct {
	Profile    string                          `json:"account"`
	Collection string                          `json:"collection"`
	Payload    *seller_api.CreatorRecomPayload `json:"payload"`
}

func NewGrabCreatorConfig() *GrabCreatorConfig {
	return &GrabCreatorConfig{
		Profile:    "testgrab@gmail.com",
		Collection: "default",
	}
}

type GrabCreator struct {
	Config   *GrabCreatorConfig
	akunRepo *repo.AccountRepo
	db       *gorm.DB
}

func NewGrabCreator(Config *GrabCreatorConfig, akunrepo *repo.AccountRepo, db *gorm.DB) *GrabCreator {
	return &GrabCreator{
		Config:   Config,
		akunRepo: akunrepo,
		db:       db,
	}
}

func (grab *GrabCreator) SaveCreator(creator *seller_api.CreatorProfile) error {
	collection := repo.Collection{
		CollName: grab.Config.Collection,
	}
	item := repo.CreatorItem{
		Collection: []*repo.Collection{
			&collection,
		},
		Nickname:    creator.Nickname.Value,
		Name:        creator.Handle.Value,
		Follower:    creator.FollowerCnt.Value,
		AverageView: creator.EcVideoAvgViewCnt.Value,
		Videos:      0,
		Region:      creator.SelectionRegion.Value,
		CreatorID:   creator.CreatorOecuid.Value,
		// Category:    creator.ProductCategories,
	}

	err := grab.db.First(&item, creator.CreatorOecuid.Value).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	item.Collection = append(item.Collection, &collection)
	err = grab.db.Save(item).Error
	return err
}

func (grab *GrabCreator) Run() error {
	log.Println("start grabbing creator")
	akun, err := grab.akunRepo.Get(grab.Config.Profile)
	if err != nil {
		return err
	}

	driver, err := driver_handler.NewDriverAccount(akun.Email, "", akun.Email, "")
	driver.Run(false, func(dctx *driver_handler.DriverContext) error {
		driver.SellerLogin(dctx)

		return err
	})
	if err != nil {
		return err
	}

	colltask, err := repo.NewCollectionTask(grab.db, grab.Config.Collection)
	defer colltask.Finish()
	if err != nil {
		return err
	}
	err = IterateCreator(context.Background(), driver, func(page int, creator *seller_api.CreatorProfile) error {
		log.Println(creator.Nickname.Value, "on page", page)
		colltask.AddCount(1)
		return grab.SaveCreator(creator)
	}, grab.Config.Payload, colltask)

	return err
}
