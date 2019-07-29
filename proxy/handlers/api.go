package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewAPIRequestRouter creates a new API request router
func NewAPIRequestRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(false)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(501)
		w.Write([]byte("not implemented"))
	})

	return r
}
