package models

import (
	"database/sql"
)

func GetAllCategories(db *sql.DB) ([]*Category , error){



}



func GetAllPostCat(db *sql.DB, user_id int) ([]*PostCat, error) {
	query :=`
	SELECT 
	P.ID , P.Title , P.content, P.datecreation
	GROUP_CONCAT(C.Name_Category , ' #') AS Categories,
	u.UserName
	FROM posts P
	JOIN
	post_category PC ON P.ID = PC.post_id
	JOIN users u ON P.user_id = u.ID
	GROUP BY P.ID
	ORDER BY P.datecreation DESC;`

	rows , err := db.Query(query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	var PostCats []*PostCat

	for rows.Next() {
		var post PostCat
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.DateCreation, &post.Categories, &post.CreatedBy)
		if err != nil {
			return nil, err
		}

		
	}	
}
