package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/path"
	"os"
)

//FilesTable display table of files returned by ioutil.ReadDir()
//TODO make each table own type, so that basePath doesn't need to be passed in
func FilesTable(files []os.FileInfo, basePath string) html.Node {

	if len(files) == 0 {
		return html.H4().Text("No files have been added")
	}

	table := html.Table().Class(flex.Grow).Children(
		html.Thead().Children(
			html.Tr().Children(
				html.Th().Text("Name"),
				html.Th().Text(""),
			),
		),

		html.Tbody().Children(
			filesTrs(files, basePath)...,
		),
	)

	return table
}

func filesTrs(files []os.FileInfo, basePath string) []html.Node {
	var nodes []html.Node
	for index := range files {
		nodes = append(nodes, fileTr(files[index], basePath))
	}
	return nodes
}

func actionsLinksForFile(file os.FileInfo, basePath string) html.Node {
	return html.Div().Children(
		html.A().Href(path.DeleteFileOrDirectoryPath(basePath + file.Name())).Text("Delete"),
	)
}

func fileTr(file os.FileInfo, basePath string) html.Node {
	var name html.Node
	if !file.IsDir() {
		name = html.Span().Text(file.Name())
	} else {
		path := path.ViewFilesPath(basePath + file.Name() + "/")
		name = html.A().Href(path).Text(file.Name())
	}

	return html.Tr().Children(
		html.Td().Children(
			name,
		),
		html.Td().Children(
			actionsLinksForFile(file, basePath),
		),
	)
}
