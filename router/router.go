package router

import (
	"log"
	"net/http"
	"time"
)

func AddHandler(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, logTimeMiddleware(handler))
}

func logTimeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}
