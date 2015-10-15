package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
)

func ManageSubtitles(video db.Video, subtitles []db.Subtitle) html.Node {
	title := "Subtitles"
	return Layout(title, html.Div().Children(
		html.H1().Text(title),

		html.Div().Children(
			html.Span().Text("Upload srt subtitle"),
			html.Form().Action(path.UploadSubtitlePath(video)).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Input().Type("file").Multiple().Name(FileParamName),
				html.Input().Type("submit"),
			),
		),
	),
	)
}
