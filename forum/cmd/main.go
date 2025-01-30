package main

import (
	"fmt"
	routes "forum/cmd/routers"
	"forum/database"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	routes.RegisterRoutes(db)

	fmt.Println("your serve on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
