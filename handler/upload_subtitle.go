package handler

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

//TODO this is almost a copy of Upload video
func UploadSubtitle(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Redirect(response, request, path.Root, http.StatusFound)
		return
	}

	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO fix assumption on buffer size
	request.ParseMultipartForm(1024 * 1024)
	form := request.MultipartForm
	formFiles := form.File[view.FormParamName]

	for _, formFile := range formFiles {
		processSubtitleFormFile(video, formFile)
	}

	http.Redirect(response, request, path.ViewVideoPath(video), http.StatusFound)
}

//TODO this is almost a copy of processVideoFromFile
func processSubtitleFormFile(video db.Video, formFile *multipart.FileHeader) {

	//TODO does this needs to be closed ?
	file, err := formFile.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received %#v", formFile.Filename)

	subtitle := db.Subtitle{Filename: formFile.Filename, VideoId: video.Id}
	db.Postgres.Save(&subtitle)
	subtitle.Save(file)

	//TODO Stop doing this as part of request
	subtitle.ConvertSrtToVtt()
}
