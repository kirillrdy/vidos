package handler

import (
	"net/http"

	"github.com/kirillrdy/vidos/path"
)

// RootHandle redirects to default route
func RootHandle(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, path.Videos.List, http.StatusFound)
}
