package view

import (
	"fmt"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

const tableClass css.Class = "videos-table"

//FormParamName is used by form for file upload and handler that processes it
const FormParamName = "file"

//Videos returns a page listing videos in a growing wrapped flexbox
func Videos(videos []db.Video) html.Node {

	var divs []html.Node

	divs = append(divs, html.Style().Text(VideoCss().String()))

	for _, video := range videos {
		divs = append(divs, Video(video))
	}

	page := html.Div().Class(hbox, grow, wrap).Children(
		divs...,
	)

	return Layout("Videos", page)
}

//VideosTable returns html table list of given videos
func VideosTable(videos []db.Video) html.Node {

	if len(videos) == 0 {
		return html.H4().Text("No videos are currently being encoded")
	}

	style := tableClass.Style(
		css.Width(size.Percent(100)),
	)

	//TODO use layout
	page := html.Div().Children(
		html.Style().Text(
			style.String(),
		),

		html.Table().Class(tableClass).Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Id"),
					html.Th().Text("File name"),
					html.Th().Text("Duration"),
					html.Th().Text("Encoded"),
					html.Th().Text("Progress"),
					html.Th(),
					html.Th(),
					html.Th(),
				),
			),

			html.Tbody().Children(
				videoTrs(videos)...,
			),
		),
	)

	return page
}

func videoTrs(videos []db.Video) []html.Node {
	var nodes []html.Node
	for index := range videos {
		nodes = append(nodes, videoTr(videos[index]))
	}
	return nodes
}

func videoTr(video db.Video) html.Node {
	return html.Tr().Children(
		html.Td().Text(video.IDString()),
		html.Td().Text(video.Filename),
		html.Td().Text(video.Duration),
		html.Td().Text(fmt.Sprint(video.Encoded)),
		html.Td().Text(fmt.Sprint(video.Progress)),
		html.Td().Children(
			html.If(video.Encoded).Then(
				html.A().Href(path.ViewVideoPath(video)).Text("View"),
			).Nodes()...,
		),

		html.Td().Children(
		//html.A().Href(path.DownloadVideoPath(video)).Text("Download"),
		),

		html.Td().Children(
		//html.A().Href(path.ReencodeVideoPath(video)).Text("Reencode"),
		),
	)

}
