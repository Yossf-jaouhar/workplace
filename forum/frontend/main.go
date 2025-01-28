package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("you server start on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmlp, err := template.ParseFiles("templete/indix.html")
	if err != nil {
		return
	}
	tmlp.Execute(w, nil)
}
