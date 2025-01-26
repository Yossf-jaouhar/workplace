package routes

import (
	"forum/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
}
