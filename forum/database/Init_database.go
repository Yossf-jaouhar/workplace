package database

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
)

func InitDatabase() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "database/forum.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	err = CreatDatabase(db)
	if err != nil {
		log.Printf("Error running migration: %v", err)
		return nil, err
	}

	return db, nil
}

func CreatDatabase(db *sql.DB) error {
	file, err := os.Open("database/sql.sql")
	if err != nil {
		return err
	}
	dataByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	dataString := string(dataByte)
	
	_, err = db.Exec(dataString)
	if err != nil {
		fmt.Println("gggggg")
		return err
	}
	return nil

}
