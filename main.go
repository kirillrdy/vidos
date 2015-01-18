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

	http.HandleFunc(path.RootPath, lib.RootHandle)
	http.HandleFunc(path.UploadPath, lib.FileUpload)
	http.HandleFunc(path.ServeFilePath, lib.ServeFile)
	http.HandleFunc(path.DownloadFilePath, lib.DownloadFile)
	http.HandleFunc(path.ReencodeFilePath, lib.ReencodeFile)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
