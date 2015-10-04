package view

import (
	"fmt"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

const TableClass css.Class = "videos-table"
const FormParamName = "file"

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

func VideosTable(videos []db.Video) html.Node {

	if len(videos) == 0 {
		return html.H4().Text("No videos are currently being encoded")
	}

	style := TableClass.Style(
		css.Width(size.Percent(100)),
	)

	//TODO use layout
	page := html.Div().Children(
		html.Style().Text(
			style.String(),
		),

		html.Table().Class(TableClass).Children(
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
