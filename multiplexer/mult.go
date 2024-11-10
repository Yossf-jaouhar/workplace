package main

import (
	"fmt"
	"net/http"
	"strings"

	"mood/funcs"
)

func main() {
	multiplexer := mux{routes: make(map[string]http.Handler)}

	multiplexer.HandleFunc("/", funcs.Home)
	multiplexer.HandleFunc("/Artist/{id}", funcs.PageArtist)
	multiplexer.HandleFunc("/search", funcs.Search)
	multiplexer.HandleFunc("/filters", funcs.FilterHandler)

	multiplexer.HandleFunc("/assets/", funcs.CssHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("127.0.0.1:8080", multiplexer)
}

// func Multipplexer(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		http.ServeFile(w, r, "Templates/index.html")
// 	case "/Artist/{id}":
// 		http.ServeFile(w, r, "Templates/info.html")
// 	default:
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

type mux struct {
	routes map[string]http.Handler
}

func (m mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.Path, "/")[1]
	handler, ok := m.routes[path]
	if ok {
		handler.ServeHTTP(w, req)
		return
	}
	handler, ok = m.routes[path+"/"]
	if ok {
		handler.ServeHTTP(w, req)
		return
	}
	w.WriteHeader(404)
}

func (m mux) HandleFunc(path string, handler http.HandlerFunc) {
	if len(path) > 1 {
		path = path[1:]
	}
	m.routes[path] = handler
}
