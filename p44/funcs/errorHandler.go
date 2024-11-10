package funcs

import (
	"html/template"
	"net/http"
)

// ErrorPageData holds the data for rendering an error page.
type ErrorPageData struct {
	Number  int
	Message string
}


// ErrorHandler handles the rendering of error pages based on the HTTP status code.
func ErrorHandler(w http.ResponseWriter, code int) {
	var message string
	switch code {
	case 500:
		message = "Internal Server Error"
		return
	case 404:
		message = "Not Found"
		return
	case 400:
		message = "Bad request"
		return
	case 405:
		message = "Method Not allowed"
		return
	}

	// for _, char := range Artis.Artists {

	// 	fmt.Println(char)

	// }

	tmpl, err := template.ParseFiles("Templates/error.html")
	if err != nil {
		http.Error(w, "Error 500", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	tmpl.Execute(w, ErrorPageData{Number: code, Message: message})
}
