package main

import (
	"fmt"
	"forum/backend"
	routes "forum/cmd/routers"
	"forum/database"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	backend.DB = db
}

func main() {
	defer backend.DB.Close()
	routes.RegisterRoutes()

	fmt.Println("your serve on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
