package models

import "database/sql"


func CountNbOfLikes(status string , postID int, db *sql.DB) int {

	query := `
	   SELECT COUNT(Status) FROM post_likes WHERE   post_id = ? AND Status = ? 
	`

	var likes int
	err := db.QueryRow(query, postID , status).Scan(likes)
	if err != nil {
		return 0
	}
	return likes
}



