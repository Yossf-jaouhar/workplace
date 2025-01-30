package handlers

import (
	"database/sql"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {

		return
	}

	tmpl, err := template.ParseFiles("frontend/templete/indix.html")
	if err != nil {

		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {

		return
	}
}
