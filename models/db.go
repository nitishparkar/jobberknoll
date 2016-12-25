package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func getDbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=nitish dbname=jobberknoll_development sslmode=disable")

	return db, err
}
