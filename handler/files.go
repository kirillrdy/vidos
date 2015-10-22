package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

//Files renderes list of files
func Files(response http.ResponseWriter, request *http.Request) {
	//TODO ModePerm possibly wrong
	//TODO Wrong place to do this
	err := os.MkdirAll(downloader.FileDir, os.ModePerm)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	basePath := request.FormValue("path")

	pathToRead := downloader.FileDir + basePath
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
