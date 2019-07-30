package handlers

import (
	"log"
	"net/http"
)

// HandleHTMLRequest returns HTML content
func HandleHTMLRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO: returning index page")

	w.WriteHeader(200)
	w.Write([]byte("hello, world"))

	// http.ServeFile(w, r, "./static/index.html")
}
