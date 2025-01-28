package handlers

import (
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, code int) {
	fmt.Println("error")
}
