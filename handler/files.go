package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/view"
)

//Files renderes list of files
func Files(response http.ResponseWriter, request *http.Request) {

	basePath := request.FormValue("path")

	pathToRead := downloader.FilesDir + basePath

	files, err := ioutil.ReadDir(pathToRead)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	view.FilesList(basePath, files).WriteTo(response)
}
