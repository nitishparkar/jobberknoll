package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func getDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", os.Getenv("JOBBERKNOLL_DB_URL"))

	return db, err
}
