package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"./handlers"
)

// Listen for HTTP traffic
func Listen(addr string) error {
	if addr == "" {
		return errors.New("addr is not defined")
	}

	r := mux.NewRouter()

	r.Handle("/", handlers.HandleWebAppRequest)

	server := &http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return server.ListenAndServe()
}
