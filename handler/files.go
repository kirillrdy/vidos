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

	files, err := ioutil.ReadDir(downloader.FileDir)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	div := html.Div().Children(
		html.H1().Text("Files"),
	)

	for _, file := range files {
		if file.IsDir() {
			div.Append(html.A().Text(file.Name()).Href(path.Files))
		} else {
			div.Append(html.Div().Children(
				html.Div().Text(file.Name()),
				html.A().Href(path.AddFileForEncodingPath(file.Name())).Text("Encode"),
			))
		}
	}

	div.Append(
		html.Span().Text("Select file to upload"),
		html.Form().Action(path.UploadFile).Attribute("enctype", "multipart/form-data").Method("POST").Children(
			html.Input().Type("file").Multiple().Name(view.FormParamName),
			html.Input().Type("submit"),
		),
	)

	view.Layout("Files", div).WriteTo(response)

}
