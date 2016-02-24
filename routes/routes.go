package routes

import (
	"github.com/kirillrdy/vidos/handler"
	"github.com/kirillrdy/vidos/path"
	"log"
	"net/http"
	"time"
)

func addHandler(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, logTimeMiddleware(handler))
}

//AddHandlers registers all the handlers with all of our middleware
func AddHandlers() {

	//Files related routes
	// addHandler(path.UploadFile, handler.UploadFile)
	// addHandler(path.Files.List, handler.Files)
	// addHandler(path.DeleteFileOrDirectory, handler.DeleteFileOrDirectory)
	// addHandler(path.AddFileForEncoding, handler.AddFileForEncoding)

	//Videos
	addHandler(path.Videos.List, handler.Videos.List)
	addHandler(path.Videos.Show, handler.Videos.Show)
	addHandler(path.Videos.Stream, handler.Stream)
	addHandler(path.Videos.Delete, handler.DeleteVideo)
	addHandler(path.Videos.New, handler.NewVideo)
	//addHandler(path.Videos.Thumbnail, handler.Thumbnail)

	//Subtitles
	//addHandler(path.ManageSubtitles, handler.ManageSubtitles)
	//addHandler(path.Subtitle, handler.Subtitle)
	// addHandler(path.UploadSubtitle, handler.SubtitlesUpload)

	// addHandler(path.Videos.Unencoded, handler.UnencodedVideos)
	//addHandler(path.Videos.Create, handler.VideosUpload)

	addHandler(path.Torrents, handler.Torrents)
	addHandler(path.TorrentStatus, handler.TorrentStatus)
	addHandler(path.AddMagnetLink, handler.AddMagnetLink)

	addHandler(path.Root, handler.RootHandle)
	http.Handle(path.Public, http.StripPrefix(path.Public, http.FileServer(http.Dir("_public"))))

}

//TODO perhaps wrap all handlers with this middleware
func logTimeMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		handler(response, request)
		log.Printf("%v %v took %v", request.Method, request.URL.Path, time.Since(start))
	}
}
