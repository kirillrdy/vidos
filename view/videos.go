package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/layout"
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
		divs = append(divs, CenterByBoxes(html.H2().Text("No videos")))
	}

	page := html.Div().Class(layout.HBox, layout.Grow, layout.Wrap).Children(
		divs...,
	)

	return Page("Videos", page)
}

func CenterByBoxes(content html.Node) html.Node {
	return html.Div().Class(layout.VBox, layout.Grow).Children(
		html.Span().Class(layout.Grow),
		html.Div().Class(layout.HBox, layout.Grow, centerItems).Children(
			html.Span().Class(layout.Grow),
			content,
			html.Span().Class(layout.Grow),
		),
		html.Span().Class(layout.Grow),
	)
}
