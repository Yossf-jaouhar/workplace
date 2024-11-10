package funcs

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"p 1")
		return
	}

	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "p 2")
		return 
	}

	FeatchArtists()

	tmpl  , err := template.ParseFiles("TT/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "p 3")
		return
	}
	             
	errr := tmpl.Execute(w, Artist)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "p 4")
		return
	}
}
