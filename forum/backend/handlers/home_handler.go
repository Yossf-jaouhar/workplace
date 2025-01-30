package handlers

import (
	"database/sql"
	"forum/backend/models"
	"forum/middleware"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed)
		return
	}

	Homepage := GetDataHomePage(r, db)

	err := RenderTemplate(w, "index.html", Homepage, http.StatusOK)
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
}

func GetDataHomePage(r *http.Request, db *sql.DB) *models.HomePage {
	HomePage := new(models.HomePage)
	token, err := middleware.VerifyCookie(r, db)

	id_user := -1

	if err == nil {
		HomePage.IsLogged = true
		ID, Name := models.GetInfos(db, token.Value)
		HomePage.UserName = Name
		totalLikes, _ := models.GetTotalLikesByUser(db, ID)
		HomePage.TotalLikes = totalLikes
		id_user = ID

	} else {
		HomePage.IsLogged = false
	}

	HomePage.PostCat, _ = models.GetAllPostCat(db, id_user)
	HomePage.Categories, _ = models.GetAllCategories(db)

	return HomePage
}
