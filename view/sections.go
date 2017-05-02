package view

import (
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web"
)

var sections = []web.Link{
	web.Link{Path: path.Videos.List, Text: "Videos"},
	web.Link{Path: path.Videos.Unencoded, Text: "Processing"},
	web.Link{Path: path.Files.List, Text: "Files"},
	web.Link{Path: path.Torrents, Text: "Torrents"},
	web.Link{Path: path.AddMagnetLink, Text: "Add Magnet link"},
}
