package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/handler"
	"github.com/kirillrdy/vidos/lib"
	"github.com/kirillrdy/vidos/path"
)

func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v taken %v", request.Method, request.URL.Path, time.Since(start))
	}
}

func main() {

	ffmpeg.CheckVersion()

	displayMemoryStats := flag.Bool("memory", false, "Print memory stats")
	port := flag.Int("port", 3001, "Port to listen on")

	if *displayMemoryStats {
		lib.StartMemoryMonitoring()
	}

	http.HandleFunc(path.Root, handler.RootHandle)
	http.HandleFunc(path.Videos, handler.Videos)
	http.HandleFunc(path.UnencodedVideos, handler.UnencodedVideos)
	http.HandleFunc(path.Upload, handler.Upload)
	http.HandleFunc(path.Serve, logMiddleware(handler.Serve))
	http.HandleFunc(path.Download, handler.Download)
	http.HandleFunc(path.Reencode, handler.ReencodeFile)
	http.HandleFunc(path.NewVideo, handler.NewVideo)

	log.Printf("Listening on %v", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
