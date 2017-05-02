package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web"
)

// RootHandle redirects to default route
func RootHandle(response http.ResponseWriter, request *http.Request) {
	web.Redirect(response, request, path.Videos.List)
}
