package handler

import (
	"github.com/kirillrdy/vidos/path"
	"net/http"
	"os"
)

func AddFileForEncoding(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO make filename a constant
	fileName := request.FormValue("filename")
	file, err := os.Open(uploadedFile{fileName}.Path())

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	processVideoFromFile(file, fileName)

	http.Redirect(response, request, path.UnencodedVideos, http.StatusFound)

}
