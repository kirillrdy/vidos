package handler

import (
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/param"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
	"log"
	"net/http"
)

func AddMagnetLink(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		view.MagnetLinkForm(response, request)
	} else {
		err := request.ParseForm()

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		magenetLink := request.Form.Get(param.MagnetLink)
		log.Print(magenetLink)

		torrentFile, err := downloader.Client.AddMagnet(magenetLink)

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		<-torrentFile.GotInfo()
		torrentFile.DownloadAll()
		http.Redirect(response, request, path.Torrents, http.StatusFound)
	}
}
