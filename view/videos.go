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

	var divs []html.Node

	divs = append(divs, html.Style().Text(VideoCSS().String()))

	for _, video := range videos {
		divs = append(divs, Video(video))
	}

	page := html.Div().Class(layout.HBox, layout.Grow, layout.Wrap).Children(
		divs...,
	)

	return Layout("Videos", page)
}
