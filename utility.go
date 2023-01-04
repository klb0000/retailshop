package retailshop

import (
	"database/sql"
	"log"
	"os"
	"path"
)

// func GetDB(file string) *gorm.DB {
// 	// check if file exist

// 	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})
// 	if err != nil {
// 		panic("failed to connect to database")
// 	}
// 	return db
// }

func GetSQLDB(file string) (*sql.DB, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return sql.Open("sqlite", path.Join(wd, "data", "data.db"))

}
