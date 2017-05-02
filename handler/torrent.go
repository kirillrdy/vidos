package handler

import (
	"bytes"
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/view"
	"net/http"
)

func Torrents(response http.ResponseWriter, request *http.Request) {
	view.TorrentsList(downloader.Client.Torrents()).WriteTo(response)
}

func TorrentStatus(response http.ResponseWriter, request *http.Request) {
	//TODO check if this needs to be new
	buffer := new(bytes.Buffer)

	downloader.Client.WriteStatus(buffer)
	view.TorrentStatus(buffer).WriteTo(response)
}
