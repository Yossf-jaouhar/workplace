package models

import (
	"database/sql"
	"log"
)

func GetInfos(db *sql.DB, token string) (int, string) {
	name := ""
	ID := 0
	err := db.QueryRow("SELECT ID, UserName FROM Users WHERE Session=?", token).Scan(&ID, &name)
	if err != nil {
		log.Printf("Error fetching user info: %v", err)
		return 0, ""
	}
	return ID, name
}
