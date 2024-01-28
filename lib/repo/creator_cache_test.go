package repo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/repo"
	"github.com/wargasipil/tiktok_rpc/scenario"
)

type Check struct {
	Val string `json:"val"`
}

func TestCacheGet(t *testing.T) {
	db := scenario.GetDatabase()

	cache := repo.NewGeneralCache(db)

	t.Run("test get cache", func(t *testing.T) {
		var vat Check
		err := cache.Get("test_key", &vat, func(key string) (interface{}, error) {
			vatt := Check{
				Val: "asdasdasd",
			}

			return vatt, nil
		})

		assert.Nil(t, err)

		err = cache.Get("test_key", &vat, func(key string) (interface{}, error) {
			vatt := Check{
				Val: "asdasdasd",
			}

			return vatt, nil
		})

		assert.Nil(t, err)
	})

	t.Run("test get error", func(t *testing.T) {
		err := cache.Clear("test_key")
		assert.Nil(t, err)

		var vat Check
		err = cache.Get("test_key", &vat, func(key string) (interface{}, error) {
			vatt := Check{
				Val: "asdasdasd",
			}

			return vatt, errors.New("gagal")
		})

		assert.NotNil(t, err)
	})

}
