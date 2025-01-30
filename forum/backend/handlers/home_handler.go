package handlers

import (
	"database/sql"
	"fmt"
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
		ID, Name := models.GetInfos(db, token.Value)
		id_user = ID

		fmt.Println(Name)
		fmt.Println(id_user)
	} else {

	}
	return HomePage
}
