package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/param"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/web"
	"net/http"
)

func MagnetLinkForm(response http.ResponseWriter, request *http.Request) {
	form := html.Form().Class(flex.VBox).Action(path.AddMagnetLink).Method("POST").Children(
		html.Input().Name(param.MagnetLink),
		html.Input().Type("submit").Value("Add"),
	)
	page := web.Page("Add Magnet Link", centerByBoxes(form))
	page.WriteTo(response)
}
