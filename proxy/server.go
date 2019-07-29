package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	handlers "./handlers"
)

// Listen for HTTP traffic
func Listen(addr string) error {
	if addr == "" {
		return errors.New("addr is not defined")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HandleHTMLRequest)

	server := &http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("SERVER: listening @", addr)

	return server.ListenAndServe()
}
