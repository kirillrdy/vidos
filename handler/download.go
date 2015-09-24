package handler

import (
	"fmt"
	"log"
	"net/http"
)

func Download(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO set content type depeding on video
	log.Printf("Serving: %v", video.FilePath())
	response.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%v"`, video.Filename))

	http.ServeFile(response, request, video.FilePath())

}
