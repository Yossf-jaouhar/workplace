package models

import (
	"database/sql"
	"forum/utils"
)

func GetAllCategories(db *sql.DB) ([]*Category , error){
	query := `SELECT ID , NameCategory FROM Category`
	rows , err := db.Query(query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	var categories []*Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.NameCategory)
		if err != nil {
			return nil , err
		}
		categories = append(categories, &category)
	}
	return categories , nil
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

		post.Like = CountNbOfLikes("like", post.ID, db)
		post.Dislike = CountNbOfLikes("dislike", post.ID, db)

		comment , err := SelectTheComment(post.ID, user_id ,db)
		if err != nil {
			return nil, err
		}
		post.Comments = append(post.Comments, comment...)
		post.NemberOfComment = len(post.Comments)

		post.Date = utils.DateFromat(post.DateCreation)
		PostCats = append(PostCats, &post)
		status := ""
		if user_id != -1 {
			status, err = GetReaction(user_id, "Post_Like", "ID_Post", post.ID, db)
			if err != nil {
				return nil, err
			}
			if status == "like" {
				post.UserLiked = true
			} else if status == "dislike" {
				post.UserDisliked = true
			}
		}	
	}
	return PostCats , nil
}
