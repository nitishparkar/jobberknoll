package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var dbConnection *gorm.DB

func ConnectToDb() (err error) {
	dbConnection, err = gorm.Open("postgres", os.Getenv("JOBBERKNOLL_DB_URL"))
	dbConnection.LogMode(true)
	return
}

func GetDbConnection() *gorm.DB {
	return dbConnection
}

func CloseDbConnection() error {
	return dbConnection.Close()
}

func MigrateDb() {
	dbConnection.AutoMigrate(&Person{})
	dbConnection.AutoMigrate(&Interaction{})
	dbConnection.Model(&Interaction{}).AddIndex("idx_interactions_person_id", "person_id")
}
