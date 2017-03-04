package main

import (
	"flag"
	"fmt"
	"github.com/kirillrdy/vidos/encoder"
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/routes"
	"log"
	"net/http"
)

func main() {

	ffmpeg.CheckVersion()

	encoder.Start()

	//TODO sort out how this is currently non blocking using buffered chan
	//db.QueueAllUnEncodedVideos()

	routes.AddHandlers()

	port := flag.Int("port", 3001, "Port to listen on")
	log.Printf("Listening on port: '%v'", *port)

	address := fmt.Sprintf(":%v", *port)

	http2Mode := false

	var err error
	if http2Mode == true {
		//TODO generate thouse for dev mode somehow
		err = http.ListenAndServeTLS(address, "localhost.cert", "localhost.key", nil)
	} else {
		err = http.ListenAndServe(address, nil)
	}

	if err != nil {
		log.Fatal(err)
	}
}
