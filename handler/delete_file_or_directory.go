package handler

import (
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"net/http"
	"os"
)

//DeleteFileOrDirectory does what name
func DeleteFileOrDirectory(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pathToRemove := request.FormValue("filepath")

	//TODO figure out why this doesn't return errors
	err = os.RemoveAll(downloader.FilesDir + pathToRemove)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, path.Files.List.String(), http.StatusFound)
}
