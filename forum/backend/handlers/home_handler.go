package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Cookies())

	tmpl, err := template.ParseFiles("frontend/templete/indix.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
