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

	// Note: ServeFile will try to detect content type based on file
	// extenttion, however mp4 is not in the list so we manually set it
	//TODO try using mime.AddExtensionType
	//TODO look into need for this, because of /etc/mime.types
	response.Header().Set("Content-Type", db.VideoMimeType)
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
