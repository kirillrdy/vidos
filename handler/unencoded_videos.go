package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/view"
)

func UnencodedVideos(response http.ResponseWriter, request *http.Request) {

	var videos []db.Video
	result := db.Session.Order("id").Where(&db.Video{Encoded: false}).Find(&videos)

	if result.Error != nil {
		log.Print(result.Error)
	}

	page := view.Videos(videos)

	io.WriteString(response, page.String())
}
