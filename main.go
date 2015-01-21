package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/lib"
	"github.com/kirillrdy/vidos/path"
)

func main() {

	ffmpeg.CheckVersion()

	displayMemoryStats := flag.Bool("memory", false, "Print memory stats")
	port := flag.Int("port", 3001, "Port to listen on")

	if *displayMemoryStats {
		lib.StartMemoryMonitoring()
	}

	http.HandleFunc(path.Root, lib.RootHandle)
	http.HandleFunc(path.Upload, lib.FileUpload)
	http.HandleFunc(path.Serve, lib.ServeFile)
	http.HandleFunc(path.Download, lib.DownloadFile)
	http.HandleFunc(path.Reencode, lib.ReencodeFile)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
