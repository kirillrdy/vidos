package main

import (
	"flag"
	"fmt"
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/routes"
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {

	ffmpeg.CheckVersion()
	//TODO sort out how this is currently non blocking using buffered chan
	//db.QueueAllUnEncodedVideos()

	routes.AddHandlers()

	port := flag.Int("port", 3001, "Port to listen on")
	log.Printf("Listening on port: '%v'", *port)
	server := http.Server{Addr: fmt.Sprintf(":%v", *port), Handler: nil}
	http2.ConfigureServer(&server, nil)
	err := server.ListenAndServeTLS("localhost.cert", "localhost.key")
	if err != nil {
		log.Fatal(err)
	}
}
