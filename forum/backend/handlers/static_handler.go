package handlers

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Stat(r.URL.Path[1:])
	if err != nil || file.IsDir() {
		HandleError(w, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, r.URL.Path[1:])
}
