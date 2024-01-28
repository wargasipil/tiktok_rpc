package repo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqliteDatabase() *gorm.DB {
	fname := "data.db"

	return CreateSqliteDatabase(fname)
}

func CreateSqliteDatabase(fname string) *gorm.DB {
	config := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Silent),
	}
	// devmode := os.Getenv("DEV_MODE") != ""
	// if !devmode {
	// 	config.Logger = logger.Default.LogMode(logger.Silent)
	// }

	db, err := gorm.Open(sqlite.Open(fname), &config)
	if err != nil {
		panic("failed to connect database")
	}
	db.Exec("VACUUM;")
	db.AutoMigrate(
		&Account{},
		&Collection{},
		&CreatorItem{},
		&AccountCacheRow{},
		&PlanCreatorTask{},
		&GeneralCacheItem{},
		&BotAuthData{},
	)
	return db
}
