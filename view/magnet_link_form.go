package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/layout"
	"github.com/kirillrdy/vidos/param"
	"github.com/kirillrdy/vidos/path"
	"net/http"
)

func MagnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Class(layout.VBox).Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(param.MagnetLink),
		html.Input().Type("submit").Value("Add"),
	)
	page := Page("Add Magnet Link", centerByBoxes(form))
	page.WriteTo(response)
}
