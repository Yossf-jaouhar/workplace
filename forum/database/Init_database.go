package database

import (
	"database/sql"
	"log"
)

func InitDatabase() *sql.DB {

	db, err := sql.Open("sqlite3", "sql.sql")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
