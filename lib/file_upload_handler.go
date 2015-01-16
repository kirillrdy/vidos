package lib

import (
	"log"
	"net/http"
)

const formParamName = "file"

const UploadPath = "/upload"
const RootPath = "/"

func FileUpload(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1024 * 1024)
	form := request.MultipartForm
	formFile := form.File[formParamName]
	//TODO handle no file submitted and only 1 file submitted
	file, err := formFile[0].Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received %#v", formFile[0].Filename)

	video := Video{Filename: formFile[0].Filename}
	video.Save(file)
	Db.Save(&video)

	http.Redirect(response, request, RootPath, http.StatusFound)
}
