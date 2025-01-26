package routes

import (
	"forum/backend/handlers"
	"net/http"
)

func RegisterRoutes() {
	
	http.HandleFunc("/", handlers.HomeHandler)

}
