package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

const fileDir = "files"

func Files(response http.ResponseWriter, request *http.Request) {
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	div := html.Div().Children(
		html.H1().Text("Files"),
	)

	for _, file := range files {
		div.Append(html.Div().Text(file.Name()))
	}

	div.Append(
		html.Span().Text("Select file to upload"),
		html.Form().Action(path.UploadFile).Attribute("enctype", "multipart/form-data").Method("POST").Children(
			html.Input().Type("file").Multiple().Name(view.FormParamName),
			html.Input().Type("submit"),
		),
	)

	io.WriteString(response, view.Layout(view.AppName, div).String())

}
