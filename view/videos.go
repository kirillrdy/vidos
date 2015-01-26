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

	page := html.Div().Class(hbox).Children(
		divs...,
	)

	//TODO maybe layout should take html.Node...
	return Layout(page)
}

func VideosTable(videos []db.Video) html.Node {

	style := TableClass.Style(
		css.Width(size.Percent(100)),
	)

	var trs []html.Node
	displayLink := func(video db.Video) string {
		if video.Encoded {
			return html.A().Href(path.ServeVideoPath(video)).Text("View").String()
		} else {
			return ""
		}
	}

	for _, video := range videos {
		tr := html.Tr().Children(
			html.Td().Text(video.IdString()),
			html.Td().Text(video.Filename),
			html.Td().Text(video.Duration),
			html.Td().Text(fmt.Sprint(video.Encoded)),
			html.Td().Text(fmt.Sprint(video.Progress)),
			html.Td().Text(
				displayLink(video),
			),

			html.Td().Children(
			//html.A().Href(path.DownloadVideoPath(video)).Text("Download"),
			),

			html.Td().Children(
			//html.A().Href(path.ReencodeVideoPath(video)).Text("Reencode"),
			),
		)
		trs = append(trs, tr)
	}

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
				trs...,
			),
		),
	)

	return Layout(page)
}
