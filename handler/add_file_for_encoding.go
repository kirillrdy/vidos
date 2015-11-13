package handler

import (
	"github.com/kirillrdy/vidos/path"
	"net/http"
	"os"
	golang_path "path"
)

//AddFileForEncoding add file for encoding from files directory
func AddFileForEncoding(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	filePath := request.FormValue(path.ParamKeys.Filepath)
	file, err := os.Open(uploadedFile{filePath}.Path())

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	processVideoFromFile(file, golang_path.Base(filePath))

	http.Redirect(response, request, path.Videos.Unencoded, http.StatusFound)

}
