package src

import (
	"net/http"
	"log"
)

type HttpRequestFunc func(w http.ResponseWriter, r *http.Request)

type HttpServer struct{}

func (hs *HttpServer) Start(addr string, onHttpRequest HttpRequestFunc) {
	http.HandleFunc("/", onHttpRequest)
	log.Fatal(http.ListenAndServe(addr, nil))
}
