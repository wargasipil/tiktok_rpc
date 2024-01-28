package repo

import (
	"sync"

	"github.com/wargasipil/tiktok_rpc/lib/seller_api"
)

type AccountCacheRow struct {
	Key string `gorm:"primaryKey;autoIncrement:false"`
	seller_api.SellerCacheData
}

type AccountCache struct {
	sync.Mutex
	Key     string
	handler func() (*seller_api.SellerCacheData, error)
	Data    *seller_api.SellerCacheData
}

func (ac *AccountCache) Get() (*seller_api.SellerCacheData, error) {

	if ac.Data != nil {
		return ac.Data, nil
	}

	ac.Lock()
	defer ac.Unlock()

	data, err := ac.handler()
	ac.Data = data

	return data, err

}
func (ac *AccountCache) SetAction(key string, handler func() (*seller_api.SellerCacheData, error)) {
	ac.Key = key
	ac.handler = handler
}
