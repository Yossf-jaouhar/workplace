package routes

import (
	"forum/backend/handlers"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/frontend/static/", handlers.StaticHandler)

	//http.HandleFunc("/addlike", middleware.Autorisation(){})
}
