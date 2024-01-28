package repo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	core_concept "github.com/wargasipil/tiktok_rpc/lib/exec_ctx"
	"github.com/wargasipil/tiktok_rpc/lib/repo"
	"github.com/wargasipil/tiktok_rpc/scenario"
)

func TestIterateRepo(t *testing.T) {
	db := repo.CreateSqliteDatabase(scenario.GetBaseTestAsset("../data.db"))

	ctxs, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ctx := core_concept.NewTaskContext(ctxs, "iterate")

	iterate, err := repo.IterateCollection(db, ctx, &repo.IterateCollConfig{
		Offset:   0,
		CollName: "default",
	})
	assert.Nil(t, err)

	found := false
	for iter := range iterate {
		t.Log(iter)
	}

	assert.True(t, found)

}
