package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/lib"
)

func main() {

	ffmpeg.CheckVersion()

	memoryFlag := flag.Bool("memory", false, "Print memory stats")

	if *memoryFlag {
		lib.StartMemoryMonitoring()
	}

	http.HandleFunc(lib.RootPath, lib.RootHandle)
	http.HandleFunc(lib.UploadPath, lib.FileUpload)
	http.HandleFunc(lib.ServeFilePath, lib.ServeFile)
	http.HandleFunc(lib.DownloadFilePath, lib.DownloadFile)
	http.HandleFunc(lib.ReencodeFilePath, lib.ReencodeFile)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
