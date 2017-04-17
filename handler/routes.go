package handler

import (
	"github.com/kirillrdy/vidos/handler/file"
	"github.com/kirillrdy/vidos/path"
	"log"
	"net/http"
	"time"
)

func AddHandler(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, logTimeMiddleware(handler))
}

//AddHandlers registers all the handlers with all of our middleware
func AddHandlers() {

	//Files related routes
	AddHandler(path.UploadFile, file.Create)
	AddHandler(path.Files.List, Files)
	// AddHandler(path.DeleteFileOrDirectory, handler.DeleteFileOrDirectory)
	// AddHandler(path.AddFileForEncoding, handler.AddFileForEncoding)

	//Videos
	AddHandler(path.Videos.List, Videos.List)
	AddHandler(path.Videos.Show, Videos.Show)
	AddHandler(path.Videos.Stream, Stream)
	AddHandler(path.Videos.Delete, DeleteVideo)
	//AddHandler(path.Videos.Thumbnail, handler.Thumbnail)

	//Subtitles
	//AddHandler(path.ManageSubtitles, handler.ManageSubtitles)
	//AddHandler(path.Subtitle, handler.Subtitle)
	// AddHandler(path.UploadSubtitle, handler.SubtitlesUpload)

	// AddHandler(path.Videos.Unencoded, handler.UnencodedVideos)

	AddHandler(path.Torrents, Torrents)
	AddHandler(path.TorrentStatus, TorrentStatus)
	AddHandler(path.AddMagnetLink, AddMagnetLink)

	AddHandler(path.Root, RootHandle)
	http.Handle(path.Public, http.StripPrefix(path.Public, http.FileServer(http.Dir("_public"))))

}

func logTimeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}
