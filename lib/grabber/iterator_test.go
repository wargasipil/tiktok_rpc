package grabber_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/api_scenario"
	"github.com/wargasipil/tiktok_rpc/lib/grabber"
	"github.com/wargasipil/tiktok_rpc/lib/repo"
	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
	"github.com/wargasipil/tiktok_rpc/scenario"
)

type LogProgress struct {
	t *testing.T
}

func (l *LogProgress) UpdateProgress(curent int, total int) error {
	l.t.Log(curent, total)
	return nil
}

func TestGrabber(t *testing.T) {
	db := scenario.GetDatabase()

	driver := api_scenario.UseDefaultDriver(t)

	payload := seller_api.NewCreatorRecomPayload()
	colltask, err := repo.NewCollectionTask(db, "test")
	assert.Nil(t, err)
	defer colltask.Finish()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	found := false
	grabber.IterateCreator(ctx, driver, func(page int, creator *seller_api.CreatorProfile) error {
		t.Log(creator)
		found = true
		cancel()
		return nil
	}, payload, &LogProgress{t: t})

	assert.True(t, found)
}
