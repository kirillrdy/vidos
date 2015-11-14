package handler

// import (
// 	"io"
// 	"net/http"

// 	"github.com/kirillrdy/vidos/db"
// 	"github.com/kirillrdy/vidos/view"
// )

// func ManageSubtitles(response http.ResponseWriter, request *http.Request) {
// 	video, err := videoFromRequest(request)

// 	if err != nil {
// 		http.Error(response, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var subtitles []db.Subtitle
// 	db.Postgres.Find(&subtitles, db.Subtitle{VideoId: video.Id})

// 	io.WriteString(response, view.ManageSubtitles(video, subtitles).String())
// }
