package main

import (
	"fmt"
	routes "forum/cmd/routers"
	"forum/database"
	"net/http"
)

func main() {

	database.InitDatabase()
	routes.RegisterRoutes()
	fmt.Println("your serve on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
