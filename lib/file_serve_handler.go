package lib

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kirillrdy/vidos/model"
)

func ServeFile(response http.ResponseWriter, request *http.Request) {
	//TODO this will be called often
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	//TODO set content type depeding on video
	response.Header().Set("Content-Type", "video/mp4")
	log.Printf("Streaming: %v", video.EncodedPath())
	http.ServeFile(response, request, video.EncodedPath())

}

func videoFromRequest(request *http.Request) (model.Video, error) {
	var video model.Video

	err := request.ParseForm()
	if err != nil {
		return video, err
	}
	idString := request.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return video, err
	}

	result := model.Db.Find(&video, id)
	if result.Error != nil {
		return video, result.Error
	}
	return video, nil
}
