package src

import (
	"log"
	"net/http"
)

type HttpRequestFunc func(w http.ResponseWriter, r *http.Request)

func StartHttpServer(addr string, onHttpRequest HttpRequestFunc) {
	http.HandleFunc("/", onHttpRequest)
	log.Fatal(http.ListenAndServe(addr, nil))
}
