package lib

import (
	"log"
	"net/http"
	"strconv"
)

const ServeFilePath = "/serve"

func ServeFile(response http.ResponseWriter, request *http.Request) {
	log.Print("Serving file")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	idString := request.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	var video Video
	result := Db.Find(&video, id)

	if result.Error != nil {
		http.Error(response, result.Error.Error(), http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "video/mp4")
	log.Printf("Trying to serve %v", video.filePath())
	http.ServeFile(response, request, video.filePath())

}
