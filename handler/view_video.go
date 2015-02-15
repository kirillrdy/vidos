package handler

import (
	"io"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/view"
)

func ViewVideo(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	var subtitles []db.Subtitle
	db.Session.Find(&subtitles, db.Subtitle{VideoId: video.Id})

	page := view.ViewVideo(video, subtitles)

	io.WriteString(response, page.String())
}
