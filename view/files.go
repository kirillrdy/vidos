package view

import (
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web/html"
	"os"
)

func FilesList(basePath string, files []os.FileInfo) html.Node {

	nodes := []html.Node{
		html.H1().Text("Files"),
		FilesTable(files, basePath),
		html.Div().Children(
			html.Span().Text("Select file to upload"),
			html.Form().Action(path.UploadFile).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Div().Children(
					html.Input().Type("file").Multiple().Name(FileParamName),
				),
				html.Div().Children(
					html.Input().Type("submit").Value("Upload"),
				),
			),
		),
	}

	return application.NewPage("Files", path.Files.List).ToHTML(nodes...)
}
