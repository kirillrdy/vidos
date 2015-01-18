package lib

import (
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/model"
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

	video := model.Video{Filename: formFile[0].Filename}
	model.Db.Save(&video)
	video.Save(file)

	video.StartEncoding()

	http.Redirect(response, request, path.Root, http.StatusFound)
}
