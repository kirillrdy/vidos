package lib

import "net/http"

const ReencodeFilePath = "/reencode"

func ReencodeFile(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	video.Reencode()

	http.Redirect(response, request, RootPath, http.StatusFound)
}
