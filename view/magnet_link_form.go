package view

import (
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/param"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web/html"
	"net/http"
)

func MagnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Class(flex.VBox).Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(param.MagnetLink),
		html.Input().Type("submit").Value("Add"),
	)
	page := application.NewPage("Add Magnet Link", path.AddMagnetLink).ToHTML(centerByBoxes(form))
	page.WriteTo(response)
}
