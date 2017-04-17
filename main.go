package main

import (
	"flag"
	"fmt"
	"github.com/kirillrdy/vidos/encoder"
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/handler"
	"log"
	"net/http"
)

func main() {
	ffmpeg.CheckVersion()
	encoder.Start()
	handler.AddHandlers()

	port := flag.Int("port", 3000, "Port to listen on")
	http2Mode := flag.Bool("http2", false, "Use http2")
	log.Printf("Listening on port: '%v'", *port)

	address := fmt.Sprintf(":%v", *port)

	var err error
	if *http2Mode == true {
		//TODO generate thouse for dev mode somehow
		err = http.ListenAndServeTLS(address, "localhost.cert", "localhost.key", nil)
	} else {
		err = http.ListenAndServe(address, nil)
	}

	if err != nil {
		log.Fatal(err)
	}
}
