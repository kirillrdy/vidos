package handler

import (
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"net/http"
	"os"
)

func DeleteFileOrDirectory(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pathToRemove := request.FormValue("filepath")

	//TODO figure out why this doesn't return errors
	err = os.RemoveAll(downloader.FileDir + pathToRemove)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, path.Files.List, http.StatusFound)
}
