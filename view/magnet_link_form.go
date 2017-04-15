package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/layout"
	"github.com/kirillrdy/vidos/path"
	"net/http"
)

var MagnetLinkFormParams = struct {
	MagnetLink string
}{
	"magnet_link",
}

//TODO make form a type that contains its param names
func MagnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Class(layout.VBox).Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(MagnetLinkFormParams.MagnetLink),
		html.Input().Type("submit").Value("Add"),
	)
	page := Page("Add Magnet Link", centerByBoxes(form))
	page.WriteTo(response)
}
