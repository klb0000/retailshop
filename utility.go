package retail_shop

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(file string) *gorm.DB {
	// check if file exist
	_, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}
