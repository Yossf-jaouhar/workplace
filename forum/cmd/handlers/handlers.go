package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}
}

func HandleError(w http.ResponseWriter, code int) {
	fmt.Println("error")
	return
}
