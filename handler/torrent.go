package handler

import (
	"bytes"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/view"
	"net/http"
)

func Torrents(response http.ResponseWriter, request *http.Request) {
	view.Layout("Torrent", view.TorrentsTable(downloader.Client.Torrents())).WriteTo(response)
}

func TorrentStatus(response http.ResponseWriter, request *http.Request) {
	buffer := new(bytes.Buffer)

	downloader.Client.WriteStatus(buffer)
	view.Layout("Torrent Status", html.Pre().Text(buffer.String())).WriteTo(response)
}
