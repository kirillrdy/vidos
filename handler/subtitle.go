package handler

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/kirillrdy/vidos/db"
// )

// //Subtitle servse a single subtitle file
// func Subtitle(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("Content-Type", "text/plain")

// 	subtitle, err := subtitleFromRequest(request)
// 	if err != nil {
// 		http.Error(response, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.ServeFile(response, request, subtitle.VttFilePath())
// }

// //TODO need to be dry, currently same as videoFromRequest
// func subtitleFromRequest(request *http.Request) (db.Subtitle, error) {
// 	var subtitle db.Subtitle

// 	idString := request.FormValue("id")
// 	id, err := strconv.Atoi(idString)
// 	if err != nil {
// 		return subtitle, err
// 	}

// 	result := db.Postgres.Find(&subtitle, id)
// 	if result.Error != nil {
// 		return subtitle, result.Error
// 	}
// 	return subtitle, nil
// }
