package handler

import (
	"net/http"
	"strconv"

	"github.com/kirillrdy/vidos/db"
)

func Subtitle(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain")

	subtitle, err := subtitleFromRequest(request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(response, request, subtitle.VttFilePath())
}

func subtitleFromRequest(request *http.Request) (db.Subtitle, error) {
	var subtitle db.Subtitle

	err := request.ParseForm()
	if err != nil {
		return subtitle, err
	}
	idString := request.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return subtitle, err
	}

	result := db.Session.Find(&subtitle, id)
	if result.Error != nil {
		return subtitle, result.Error
	}
	return subtitle, nil
}
