package handler

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
	"net/http"
)

const MagnetLink = "magnet_link"

func AddMagnetLink(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		magnetLinkForm(response, request)
	}
}

func magnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(MagnetLink),
	)
	page := view.Layout("Add Magnet Link", form)
	page.WriteTo(response)
}
