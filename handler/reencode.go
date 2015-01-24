package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/path"
)

func ReencodeFile(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	video.Reencode()

	http.Redirect(response, request, path.Root, http.StatusFound)
}
