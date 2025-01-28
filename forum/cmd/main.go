package main

import (
	"fmt"
	"forum/backend"
	routes "forum/cmd/routers"
	"forum/database"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	backend.DB = database.InitDatabase()
	routes.RegisterRoutes()
	fmt.Println("your serve on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
