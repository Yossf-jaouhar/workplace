package funcs

import (
	"net/http"
	"os"
)

// CssHandler handles requests for CSS files in the "assets" directory.
func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		files, err := os.ReadDir("static")
		if err != nil {
			return
		}
		for _, file := range files {
			if r.URL.Path == "/static/"+file.Name() {

				fs := http.Dir("static")
				http.StripPrefix("/static", http.FileServer(fs)).ServeHTTP(w, r)
				return
			}
		}
		
		return
	} else {
		return
	}
}
