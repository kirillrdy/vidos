package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kirillrdy/vidos/db"
)

func Serve(response http.ResponseWriter, request *http.Request) {

	log.Print(request.Header.Get("Range"))

	//TODO this will be called often
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO set content type depeding on video
	response.Header().Set("Content-Type", "video/mp4")
	log.Printf("Streaming: %v", video.EncodedPath())
	http.ServeFile(response, request, video.EncodedPath())
	log.Printf("Finished streaming: %v", video.EncodedPath())
}

func videoFromRequest(request *http.Request) (db.Video, error) {
	var video db.Video

	err := request.ParseForm()
	if err != nil {
		return video, err
	}
	idString := request.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return video, err
	}

	result := db.Postgres.Find(&video, id)
	if result.Error != nil {
		return video, result.Error
	}
	return video, nil
}
