package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
	"os"
)

//FilesTable display table of files returned by ioutil.ReadDir()
func FilesTable(files []os.FileInfo, basePath string) html.Node {

	if len(files) == 0 {
		return html.H4().Text("No files have been added")
	}

	style := TableClass.Style(
		css.Width(size.Percent(100)),
	)

	//TODO use layout
	page := html.Div().Children(
		html.Style().Text(
			style.String(),
		),

		html.Table().Class(TableClass).Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Name"),
					html.Th().Text(""),
				),
			),

			html.Tbody().Children(
				filesTrs(files, basePath)...,
			),
		),
	)

	return page
}

func filesTrs(files []os.FileInfo, basePath string) []html.Node {
	var nodes []html.Node
	for index := range files {
		nodes = append(nodes, fileTr(files[index], basePath))
	}
	return nodes
}

func fileTr(file os.FileInfo, basePath string) html.Node {
	var link html.Node
	if !file.IsDir() {
		link = html.A().Href(path.AddFileForEncodingPath(basePath + file.Name())).Text("Encode")
	} else {
		link = html.A().Href(path.ViewFilesPath(basePath + file.Name() + "/")).Text("View")
	}

	return html.Tr().Children(
		html.Td().Text(file.Name()),
		html.Td().Children(
			link,
		),
	)
}
