package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

//Files renderes list of files
func Files(response http.ResponseWriter, request *http.Request) {

	basePath := request.FormValue("path")

	pathToRead := downloader.FilesDir + basePath
	log.Printf("%#v", pathToRead)

	files, err := ioutil.ReadDir(pathToRead)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	nodes := []html.Node{
		html.H1().Text("Files"),
		view.FilesTable(files, basePath),
		html.Div().Children(
			html.Span().Text("Select file to upload"),
			html.Form().Action(path.UploadFile).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Div().Children(
					html.Input().Type("file").Multiple().Name(view.FileParamName),
				),
				html.Div().Children(
					html.Input().Type("submit").Value("Upload"),
				),
			),
		),
	}

	view.Layout("Files", nodes...).WriteTo(response)

}
