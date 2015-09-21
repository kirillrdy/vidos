package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/handler"
	"github.com/kirillrdy/vidos/path"
)

//TODO move to own package
func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}

func main() {

	ffmpeg.CheckVersion()
	db.QueueAllUnEncodedVideos()

	//TODO move flags somewhere
	port := flag.Int("port", 3001, "Port to listen on")

	//TODO Getting more routes move somewhere
	http.HandleFunc(path.UploadFile, handler.UploadFile)
	http.HandleFunc(path.Files, handler.Files)
	http.HandleFunc(path.ViewVideo, handler.ViewVideo)
	http.HandleFunc(path.Subtitle, handler.Subtitle)
	http.HandleFunc(path.Videos, handler.Videos)
	http.HandleFunc(path.UnencodedVideos, handler.UnencodedVideos)
	http.HandleFunc(path.Upload, handler.Upload)
	http.HandleFunc(path.UploadSubtitle, handler.UploadSubtitle)
	http.HandleFunc(path.Serve, logMiddleware(handler.Serve))
	http.HandleFunc(path.Download, handler.Download)
	http.HandleFunc(path.Reencode, handler.ReencodeFile)
	http.HandleFunc(path.Delete, handler.DeleteVideo)
	http.HandleFunc(path.NewVideo, handler.NewVideo)
	http.HandleFunc(path.Thumbnail, handler.Thumbnail)
	http.HandleFunc(path.ManageSubtitles, handler.ManageSubtitles)

	http.HandleFunc(path.Root, handler.RootHandle)
	http.Handle(path.Public, http.StripPrefix(path.Public, http.FileServer(http.Dir("public"))))

	log.Printf("Listening on port: '%v'", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
