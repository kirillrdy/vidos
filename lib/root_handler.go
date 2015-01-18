package lib

import (
	"io"
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/model"
	"github.com/kirillrdy/vidos/view"
)

func RootHandle(response http.ResponseWriter, request *http.Request) {

	var videos []model.Video
	result := model.Db.Order("id").Find(&videos)
	if result.Error != nil {
		log.Print(result.Error)
	}

	page := view.Videos(videos)

	io.WriteString(response, page.String())
}
