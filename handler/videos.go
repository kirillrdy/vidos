package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/view"
)

//Videos contains all handlers realted to videos
var Videos = struct {
	List func(response http.ResponseWriter, request *http.Request)
	Show func(response http.ResponseWriter, request *http.Request)
}{
	//List
	func(response http.ResponseWriter, request *http.Request) {
		var videos []db.Video
		result := db.Postgres.Order("id desc").Where(&db.Video{Encoded: true}).Find(&videos)
		//TODO create a function that will do this, also possibly with better layout
		if result.Error != nil {
			http.Error(response, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		view.Videos(videos).WriteTo(response)
	},

	//Show
	func(response http.ResponseWriter, request *http.Request) {
		video, err := videoFromRequest(request)

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var subtitles []db.Subtitle
		result := db.Postgres.Find(&subtitles, db.Subtitle{VideoId: video.Id})
		if result.Error != nil {
			http.Error(response, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		view.Layout(video.Filename, view.ViewVideo(video, subtitles)).WriteTo(response)
	},
}
