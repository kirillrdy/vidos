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
func logTimeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}

func addHandlers() {
	//TODO Getting more routes move somewhere

	//Files related routes
	http.HandleFunc(path.UploadFile, handler.UploadFile)
	http.HandleFunc(path.Files.List, handler.Files)
	http.HandleFunc(path.DeleteFileOrDirectory, handler.DeleteFileOrDirectory)
	http.HandleFunc(path.AddFileForEncoding, handler.AddFileForEncoding)

	//Videos
	http.HandleFunc(path.Videos.List, handler.Videos.List)
	http.HandleFunc(path.Videos.Show, handler.Videos.Show)
	http.HandleFunc(path.Videos.Stream, logTimeMiddleware(handler.Stream))
	http.HandleFunc(path.Videos.Download, handler.Download)
	http.HandleFunc(path.Videos.Delete, handler.DeleteVideo)
	http.HandleFunc(path.Videos.New, handler.NewVideo)
	http.HandleFunc(path.Videos.Thumbnail, handler.Thumbnail)

	//Subtitles
	http.HandleFunc(path.ManageSubtitles, handler.ManageSubtitles)
	http.HandleFunc(path.Subtitle, handler.Subtitle)
	http.HandleFunc(path.UploadSubtitle, handler.SubtitlesUpload)

	http.HandleFunc(path.Videos.Unencoded, handler.UnencodedVideos)
	http.HandleFunc(path.Videos.Create, handler.VideosUpload)

	http.HandleFunc(path.Torrents, handler.Torrents)
	http.HandleFunc(path.TorrentStatus, handler.TorrentStatus)
	http.HandleFunc(path.AddMagnetLink, handler.AddMagnetLink)

	http.HandleFunc(path.Root, handler.RootHandle)
	http.Handle(path.Public, http.StripPrefix(path.Public, http.FileServer(http.Dir("public"))))

}

func main() {

	ffmpeg.CheckVersion()
	db.QueueAllUnEncodedVideos()

	//TODO move flags somewhere
	port := flag.Int("port", 3001, "Port to listen on")

	addHandlers()

	log.Printf("Listening on port: '%v'", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
