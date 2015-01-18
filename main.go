package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/lib"
	"github.com/kirillrdy/vidos/path"
)

func main() {

	ffmpeg.CheckVersion()

	memoryFlag := flag.Bool("memory", false, "Print memory stats")

	if *memoryFlag {
		lib.StartMemoryMonitoring()
	}

	http.HandleFunc(path.Root, lib.RootHandle)
	http.HandleFunc(path.Upload, lib.FileUpload)
	http.HandleFunc(path.Serve, lib.ServeFile)
	http.HandleFunc(path.Download, lib.DownloadFile)
	http.HandleFunc(path.Reencode, lib.ReencodeFile)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
