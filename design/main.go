package main

import (
	"fmt"
	"net/http"
	"web/funcs"
)

func main() {
	http.HandleFunc("/", funcs.Home)

	http.HandleFunc("/static/", funcs.CssHandler)

	fmt.Printf("Your server is running on %s\n", "http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
