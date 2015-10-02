package handler

import (
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func NewVideo(response http.ResponseWriter, request *http.Request) {
	page := view.Layout("Upload new video",
		html.Div().Children(
			html.Span().Text("Select file to upload"),
			html.Form().Action(path.Videos.Create).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Input().Type("file").Multiple().Name(view.FormParamName),
				html.Input().Type("submit").Value("Upload"),
			),
		),
	)
	page.WriteTo(response)
}
