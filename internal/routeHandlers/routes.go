// Package routehandlers handles all the requests comming to the app
package routehandlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRoutes handles all the endpoints and related handler function for them
func HandleRoutes(router *mux.Router) {
	router.HandleFunc("/", homeHandler)

	fs := http.FileServer(http.Dir("./template/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

// homeHandler handles the request to the "/" endpoint
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Home Handler!"))
}
