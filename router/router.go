package router

import (
	"github.com/kirillrdy/web"
	"log"
	"net/http"
	"time"
)

// AddHandler adds a handle function for a given path
func AddHandler(path web.Path, handler http.HandlerFunc) {
	http.HandleFunc(path.String(), logTimeMiddleware(handler))
}

func logTimeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}
