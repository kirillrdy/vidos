package handler

import (
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/view"
)

func VideosList(response http.ResponseWriter, request *http.Request) {

	var videos []db.Video
	result := db.Postgres.Order("id desc").Where(&db.Video{Encoded: true}).Find(&videos)

	if result.Error != nil {
		log.Print(result.Error)
	}

	page := view.Videos(videos)
	page.WriteTo(response)
}
