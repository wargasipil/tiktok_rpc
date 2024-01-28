package repo

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type GeneralCache struct {
	db *gorm.DB
}

func NewGeneralCache(db *gorm.DB) *GeneralCache {
	return &GeneralCache{
		db: db,
	}
}

type GeneralCacheItem struct {
	Key  string `gorm:"primaryKey;autoIncrement:false"`
	Data []byte
}

func (cache *GeneralCache) Clear(key string) error {
	item := GeneralCacheItem{
		Key: key,
	}

	return cache.db.Delete(&item).Error
}

func (cache *GeneralCache) Get(key string, hasil interface{}, handler func(key string) (interface{}, error)) error {
	item := GeneralCacheItem{
		Key: key,
	}
	err := cache.db.First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		data, err := handler(key)
		if err != nil {
			return err
		}
		databyte, err := json.Marshal(data)
		if err != nil {
			return err
		}
		item.Data = databyte

		err = cache.db.Save(&item).Error
		if err != nil {
			return err
		}
	} else {
		return err
	}
	err = json.Unmarshal(item.Data, hasil)
	if err != nil {
		return err
	}

	return nil
}
