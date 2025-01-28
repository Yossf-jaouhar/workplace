package handlers

import (
	"fmt"
	"forum/backend"
	"net/http"
	"text/template"
)

func HandleError(w http.ResponseWriter, code int) {
	tmpl, err := template.ParseFiles("frontend/templete/error.html")
	if err != nil {
		fmt.Println("error")
		return
	}

	status := http.StatusText(code)

	var data backend.Err
	data.Code = code
	data.Type = status

	w.WriteHeader(code)

	tmpl.Execute(w, data)
}
