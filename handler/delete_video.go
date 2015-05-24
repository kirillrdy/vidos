package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/path"
)

func DeleteVideo(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	video.Delete()
	http.Redirect(response, request, path.Root, http.StatusFound)
}
