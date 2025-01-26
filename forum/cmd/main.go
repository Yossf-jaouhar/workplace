package main

import (
	"fmt"
	"net/http"
)

func main() {

	//InitDatabase(dbPath string)
	RegisterRoutes()
	fmt.Println("your serve on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
