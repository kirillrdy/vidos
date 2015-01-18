package lib

import (
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func FileUpload(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1024 * 1024)
	form := request.MultipartForm
	formFile := form.File[view.FormParamName]
	//TODO handle no file submitted and only 1 file submitted
	file, err := formFile[0].Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received %#v", formFile[0].Filename)

	video := db.Video{Filename: formFile[0].Filename}
	db.Session.Save(&video)
	video.Save(file)
	video.CalculateDuration()
	video.StartEncoding()

	http.Redirect(response, request, path.Root, http.StatusFound)
}
