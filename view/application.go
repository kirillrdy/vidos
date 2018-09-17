package view

import (
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/web"
)

var application = web.Application{Name: appName, Menu: []web.Page{
	{Title: "Videos", Path: path.Videos.List},
	{Title: "Processing", Path: path.Videos.Unencoded},
	{Title: "Files", Path: path.Files.List},
	{Title: "Torrents", Path: path.Torrents},
	{Title: "Add Magnet Link", Path: path.AddMagnetLink},
}}

const appName = "Видос"
