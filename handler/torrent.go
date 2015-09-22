package handler

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/view"
	"net/http"
)

func Torrents(response http.ResponseWriter, request *http.Request) {
	view.Layout("Torrent", html.H1().Text("Torrents")).WriteTo(response)
}
