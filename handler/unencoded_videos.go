package handler

import (
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/view"
)

func UnencodedVideos(response http.ResponseWriter, request *http.Request) {

	var videos []db.Video
	result := db.Postgres.Order("id").Not(&db.Video{Encoded: true}).Find(&videos)

	if result.Error != nil {
		log.Print(result.Error)
	}

	page := view.VideosTable(videos)

	view.Layout("Videos", page).WriteTo(response)
}
