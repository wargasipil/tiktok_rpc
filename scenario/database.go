package scenario

import (
	"github.com/wargasipil/tiktok_rpc/lib/repo"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	return repo.CreateSqliteDatabase(GetBaseTestAsset("database_test.db"))
}
