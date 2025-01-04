package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var DB *sql.DB

func InitDatabase(filepath string) {
	var err error

	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Database connected successfully!")

	// Create tables
	createTables()
}

func createTables() {
	for _, query := range Queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Failed to create table:", err)
		}
	}

	log.Println("All tables created successfully!")
}
