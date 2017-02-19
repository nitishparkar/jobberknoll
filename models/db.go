package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=nitish dbname=jobberknoll_development sslmode=disable")

	return db, err
}
