package routes

import (
	"database/sql"
	"forum/backend/handlers"
	user "forum/backend/handlers/authh"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, db)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		user.Login(w, r, db)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})

	
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		user.Logout(w, r, db)
	})

	http.HandleFunc("/frontend/static/", handlers.StaticHandler)

	//http.HandleFunc("/addlike", middleware.Autorisation())
}
