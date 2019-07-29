package handlers

import "net/http"

// HandleWebAppRequest returns the HTML for the webapp
func HandleWebAppRequest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
