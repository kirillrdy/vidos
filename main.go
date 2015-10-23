package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/routes"
)

func main() {

	ffmpeg.CheckVersion()
	//TODO sort out how this is currently non blocking using buffered chan
	db.QueueAllUnEncodedVideos()

	routes.AddHandlers()

	port := flag.Int("port", 3001, "Port to listen on")
	log.Printf("Listening on port: '%v'", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
