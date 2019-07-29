package server

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"

	handlers "./handlers"
	targets "./targets"
)

var targetCollection = &targets.TargetCollection{}
var httpClient = &http.Client{}

// Listen for HTTP traffic
func Listen(addr, cpanelHostname string) error {
	// todo: remove this later
	targetCollection.SetTarget("localhost:8080", &targets.SimpleTarget{"symmtric.solutions", "https"})

	if addr == "" {
		return errors.New("addr is not defined")
	}

	if cpanelHostname == "" {
		return errors.New("cpanelHostname is not defined")
	}

	r := mux.NewRouter()

	// CPANEL stuff
	r.Host(cpanelHostname).Path("/").HandlerFunc(handlers.HandleHTMLRequest)
	r.Host(cpanelHostname).PathPrefix("/api").PathPrefix("/api").Handler(handlers.NewAPIRequestRouter())
	r.Host(cpanelHostname).PathPrefix("/intercom").HandlerFunc(handlers.HandleSocketRequest)

	// proxy stuff
	r.HandleFunc("/", proxyHandlerFunc)

	server := &http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("SERVER: listening @", addr)

	return server.ListenAndServe()
}

func proxyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	t := targetCollection.GetTarget(r.Host)

	if t == nil {
		w.WriteHeader(404)
		w.Write([]byte("target not found for that hostname"))
		return
	}

	u := url.URL{
		Scheme: t.Scheme(),
		Host:   t.Host(),
		Path:   r.RequestURI,
	}

	pReq, err := http.NewRequest(r.Method, u.String(), r.Body)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("an error occured while creating the proxy request"))
		return
	}

	pResp, err := httpClient.Do(pReq)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("an error occured while getting the proxy response"))
		return
	}

	defer pResp.Body.Close()

	for key, values := range pResp.Header {
		w.Header().Set(key, strings.Join(values, "; "))
	}

	w.WriteHeader(pResp.StatusCode)

	if pResp.ContentLength != 0 {
		buffer := make([]byte, 128*1024)

		for {
			i, err := pResp.Body.Read(buffer)

			if err != nil {
				if err == io.EOF {
					w.Write(buffer[:i])
				} else {
					log.Println("ERROR: ", err)
				}

				break
			}

			w.Write(buffer)
		}
	}
}
