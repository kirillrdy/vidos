package view

import (
	"fmt"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/model"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

const TableClass css.Class = "videos-table"
const FormParamName = "file"

func Videos(videos []model.Video) html.Node {

	style := TableClass.Style(
		css.Margin(size.Auto),
	)

	var trs []html.Node
	displayLink := func(video model.Video) string {
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
			html.Td().Text(fmt.Sprint(video.Encoded)),
			html.Td().Text(fmt.Sprint(video.Progress)),
			html.Td().Text(
				displayLink(video),
			),

			html.Td().Children(
				html.A().Href(path.DownloadVideoPath(video)).Text("Download"),
			),

			html.Td().Children(
				html.A().Href(path.ReencodeVideoPath(video)).Text("Reencode"),
			),
		)
		trs = append(trs, tr)
	}

	page := html.Html().Children(
		html.Style().Text(
			style.String(),
		),
		html.Table().Class(TableClass).Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Id"),
					html.Th().Text("File name"),
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
		html.Form().Action(path.UploadPath).Attribute("enctype", "multipart/form-data").Method("POST").Children(
			html.Input().Type("file").Name(FormParamName),
			html.Input().Type("submit"),
		),
	)

	return page
}
