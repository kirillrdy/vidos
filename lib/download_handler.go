package lib

import (
	"fmt"
	"log"
	"net/http"
)

const DownloadFilePath = "/download"

func DownloadFile(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	//TODO set content type depeding on video
	log.Printf("Serving: %v", video.filePath())
	response.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%v"`, video.Filename))

	http.ServeFile(response, request, video.filePath())

}
