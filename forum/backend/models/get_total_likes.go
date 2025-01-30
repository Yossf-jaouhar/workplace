package models

import "database/sql"


func GetTotalLikesByUser(db *sql.DB , id int) (int , error) {

	var TotalLikes int
	err := db.QueryRow("SELECT COUNT(Status) FROM post_likes WHERE   user_id = ?", id).Scan(TotalLikes)
	if err != nil {
		return 0, err
	}
	return TotalLikes , nil
}