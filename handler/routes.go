package handler

import (
	"github.com/kirillrdy/vidos/handler/file"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/router"
	"github.com/kirillrdy/web/css"
	"log"
	"net/http"
	"time"
)

//AddHandlers registers all the handlers with all of our middleware
func AddHandlers() {

	//Files related routes
	router.AddHandler(path.UploadFile, file.Create)
	router.AddHandler(path.Files.List, Files)
	// router.AddHandler(path.DeleteFileOrDirectory, handler.DeleteFileOrDirectory)
	// router.AddHandler(path.AddFileForEncoding, handler.AddFileForEncoding)

	//Videos
	router.AddHandler(path.Videos.List, Videos.List)
	router.AddHandler(path.Videos.Show, Videos.Show)
	router.AddHandler(path.Videos.Stream, Stream)
	router.AddHandler(path.Videos.Delete, DeleteVideo)
	//router.AddHandler(path.Videos.Thumbnail, handler.Thumbnail)

	//Subtitles
	//router.AddHandler(path.ManageSubtitles, handler.ManageSubtitles)
	//router.AddHandler(path.Subtitle, handler.Subtitle)
	// router.AddHandler(path.UploadSubtitle, handler.SubtitlesUpload)

	// router.AddHandler(path.Videos.Unencoded, handler.UnencodedVideos)

	router.AddHandler(path.Torrents, Torrents)
	router.AddHandler(path.TorrentStatus, TorrentStatus)
	router.AddHandler(path.AddMagnetLink, AddMagnetLink)

	router.AddHandler(path.Root, RootHandle)
	router.AddHandler(path.CSSReset, css.ServeResetCSS)

}

func logTimeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}
