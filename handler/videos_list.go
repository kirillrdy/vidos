package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/view"
	"github.com/kirillrdy/vidos/web"
)

//Videos contains all handlers realted to videos
var Videos = struct {
	List func(response http.ResponseWriter, request *http.Request)
	Show func(response http.ResponseWriter, request *http.Request)
}{
	//List
	func(response http.ResponseWriter, request *http.Request) {
		videos, err := fs.Videos()

		//TODO create a function that will do this, also possibly with better layout
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		view.Videos(videos).WriteTo(response)
	},

	//Show
	func(response http.ResponseWriter, request *http.Request) {
		video := videoFromRequest(request)

		// var subtitles []db.Subtitle
		// result := db.Postgres.Find(&subtitles, db.Subtitle{VideoId: video.Id})
		// if result.Error != nil {
		// 	http.Error(response, result.Error.Error(), http.StatusInternalServerError)
		// 	return
		// }

		web.Page(video.Filename(), view.VideoShowPage(video)).WriteTo(response)
	},
}
