package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/fs"
	"github.com/sparkymat/webdsl/css"
)

const tableClass css.Class = "videos-table"

//FileParamName is used by form for file upload and handler that processes it
const FileParamName = "file"

//Videos returns a page listing videos in a growing wrapped flexbox
func Videos(videos []fs.Video) html.Node {

	divs := make([]html.Node, 0, len(videos))

	divs = append(divs, html.Style().Text(VideoCSS().String()))

	for _, video := range videos {
		divs = append(divs, Video(video))
	}

	if len(videos) == 0 {
		divs = append(divs, centerByBoxes(html.H2().Text("No videos")))
	}

	page := html.Div().Class(flex.HBox, flex.Grow, flex.Wrap).Children(
		divs...,
	)

	return Page("Videos", page)
}

func centerByBoxes(content html.Node) html.Node {
	return html.Div().Class(flex.VBox, flex.Grow).Children(
		html.Span().Class(flex.Grow),
		html.Div().Class(flex.HBox, flex.Grow, centerItems).Children(
			html.Span().Class(flex.Grow),
			content,
			html.Span().Class(flex.Grow),
		),
		html.Span().Class(flex.Grow),
	)
}
