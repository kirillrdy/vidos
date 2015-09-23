package handler

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
	"log"
	"net/http"
)

const MagnetLink = "magnet_link"

func AddMagnetLink(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		magnetLinkForm(response, request)
	} else {
		request.ParseForm()
		log.Print(request.Form.Get(MagnetLink))

		torrentFile, err := downloader.Client.AddMagnet(request.Form.Get(MagnetLink))
		//torrentFile, err := downloader.Client.AddMagnet("magnet:?xt=urn:btih:ZOCMZQIPFFW7OLLMIC5HUB6BPCSDEOQU")

		//Return errors to http client
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}

		<-torrentFile.GotInfo()
		torrentFile.DownloadAll()
		http.Redirect(response, request, path.Torrents, http.StatusFound)
	}
}

func magnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(MagnetLink),
		html.Input().Type("submit"),
	)
	page := view.Layout("Add Magnet Link", form)
	page.WriteTo(response)
}
