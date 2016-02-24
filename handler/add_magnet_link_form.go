package handler

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
	"log"
	"net/http"
)

const magnetLinkParamName = "magnet_link"

func AddMagnetLink(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		magnetLinkForm(response, request)
	} else {
		request.ParseForm()
		log.Print(request.Form.Get(magnetLinkParamName))

		torrentFile, err := downloader.Client.AddMagnet(request.Form.Get(magnetLinkParamName))

		//Return errors to http client
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		<-torrentFile.GotInfo()
		torrentFile.DownloadAll()
		http.Redirect(response, request, path.Torrents, http.StatusFound)
	}
}

func magnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(magnetLinkParamName),
		html.Input().Type("submit").Value("Add"),
	)
	page := view.Layout("Add Magnet Link", form)
	page.WriteTo(response)
}
