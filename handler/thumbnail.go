package handler

import "net/http"

func Thumbnail(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	http.ServeFile(response, request, video.ThumbnailPath())
}
