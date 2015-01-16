package lib

import (
	"io"
	"log"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
)

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
		)
		trs = append(trs, tr)
	}

	page := html.Html().Children(
		html.Table().Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Id"),
					html.Th().Text("File name"),
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
