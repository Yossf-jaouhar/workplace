package main

import (
	"fmt"
	"net/http"

	"ppppp/funcs"
)

func main() {
	http.HandleFunc("/", funcs.Home)
	fmt.Println("server in http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
