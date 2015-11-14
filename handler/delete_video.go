package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/path"
)

//DeleteVideo removes a video file from filesystem fs
func DeleteVideo(response http.ResponseWriter, request *http.Request) {
	video := videoFromRequest(request)

	err := video.Delete()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, path.Root, http.StatusFound)
}
