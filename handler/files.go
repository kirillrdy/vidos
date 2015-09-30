package handler

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func Files(response http.ResponseWriter, request *http.Request) {
	//TODO ModePerm possibly wrong
	//TODO Wrong place to do this
	err := os.MkdirAll(downloader.FileDir, os.ModePerm)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	err = request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	basePath := request.FormValue("path")

	files, err := ioutil.ReadDir(downloader.FileDir + basePath)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	div := html.Div().Children(
		html.H1().Text("Files"),
		html.Div().Children(view.FilesTable(files, basePath)),
		html.Div().Children(
			html.Span().Text("Select file to upload"),
			html.Form().Action(path.UploadFile).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Div().Children(
					html.Input().Type("file").Multiple().Name(view.FormParamName),
				),
				html.Div().Children(
					html.Input().Type("submit").Value("Upload"),
				),
			),
		),
	)

	view.Layout("Files", div).WriteTo(response)

}
