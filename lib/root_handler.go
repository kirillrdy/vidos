package lib

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
)

func ServeVideoPath(video Video) string {
	return fmt.Sprintf("%v?id=%v", ServeFilePath, video.Id)
}

func DownloadVideoPath(video Video) string {
	return fmt.Sprintf("%v?id=%v", DownloadFilePath, video.Id)
}

func RootHandle(response http.ResponseWriter, request *http.Request) {

	var trs []html.Node
	var videos []Video
	result := Db.Find(&videos)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	for _, video := range videos {
		tr := html.Tr().Children(
			html.Td().Text(video.IdString()),
			html.Td().Text(video.Filename),
			html.Td().Children(
				html.A().Href(ServeVideoPath(video)).Text("View"),
			),

			html.Td().Children(
				html.A().Href(DownloadVideoPath(video)).Text("Download"),
			),
		)
		trs = append(trs, tr)
	}

	page := html.Html().Children(
		html.Table().Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Id"),
					html.Th().Text("File name"),
					html.Th(),
					html.Th(),
				),
			),

			html.Tbody().Children(
				trs...,
			),
		),
		html.Form().Action(UploadPath).Attribute("enctype", "multipart/form-data").Method("POST").Children(
			html.Input().Type("file").Name(formParamName),
			html.Input().Type("submit"),
		),
	)
	io.WriteString(response, page.String())
}
