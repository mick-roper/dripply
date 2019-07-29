package handlers

import "net/http"

// HandleHTMLRequest returns HTML content
func HandleHTMLRequest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
