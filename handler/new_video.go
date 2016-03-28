package handler

import (
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/layout"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func NewVideo(response http.ResponseWriter, request *http.Request) {
	page := view.Layout("Upload new video",
		view.CenterByBoxes(html.Div().Class(layout.VBox).Children(
			html.Span().Text("Select file to upload"),
			html.Form().Action(path.Videos.Create).Class(layout.VBox).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Input().Type("file").Multiple().Name(view.FileParamName),
				html.Input().Type("submit").Value("Upload"),
			),
		),
		))
	page.WriteTo(response)
}
