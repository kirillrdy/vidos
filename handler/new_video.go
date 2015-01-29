package handler

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func NewVideo(response http.ResponseWriter, request *http.Request) {
	page := view.Layout(
		html.Div().Children(
			html.Span().Text("Select file to upload"),
			html.Form().Action(path.Upload).Multiple().Method("POST").Children(
				html.Input().Type("file").Attribute("multiple", "multiple").Name(view.FormParamName),
				html.Input().Type("submit"),
			),
		),
	)
	io.WriteString(response, page.String())
}
