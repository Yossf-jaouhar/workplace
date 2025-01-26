package main

import (
	"fmt"
	routes "forum/routers"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("your serve on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
