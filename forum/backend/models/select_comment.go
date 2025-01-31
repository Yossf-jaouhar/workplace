package models

import (
	"database/sql"
	"forum/utils"
	"log"
	"strconv"
)

func SelectTheComment(postid int, userid int, db *sql.DB) ([]*Comment, error) {
	query := `
	  SELECT C.ID, C.Content, U.UserName, C.datecreation
	  FROM Comment C
	  JOIN posts P ON P.ID = C.post_id
	  JOIN users U ON U.ID = C.user_id
	  WHERE post_id = ?
	`

	rows, err := db.Query(query, postid)
	if err != nil {
		return nil, err
	}

	var Comments []*Comment
	for rows.Next() {
		var eachComment Comment

		err := rows.Scan(&eachComment.ID, &eachComment.Content, &eachComment.CreatedBy, &eachComment.DateCreation)
		if err != nil {
			return nil, err
		}
		eachComment.Date = utils.DateFromat(eachComment.DateCreation)

		eachComment.Like, err = NbOfReactInComct(eachComment.ID, db, "like")
		eachComment.Dislike, err = NbOfReactInComct(eachComment.ID, db, "dislike")
		if err != nil {
			log.Fatal(err)
		}
		status := ""
		if userid != -1 {
			id, _ := strconv.Atoi(eachComment.ID)
			status, err = GetReaction(userid, "Comment_Like", "ID_Comment", id, db)
			if err != nil {
				return nil, err
			}
			if status == "like" {
				eachComment.CUserLiked = true
			} else if status == "dislike" {
				eachComment.CUserDisliked = true
			}
		}
		Comments = append(Comments, &eachComment)
	}
	return Comments , nil
}
