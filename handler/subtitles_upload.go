package handler

import (
	"io"
	"net/http"

	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
)

//SubtitlesUpload handles uploading of subtitles
func SubtitlesUpload(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Redirect(response, request, path.Root, http.StatusFound)
		return
	}

	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	handleMultiFileUpload(response, request, func(file io.ReadCloser, fileName string) {
		defer file.Close()

		subtitle := db.Subtitle{Filename: fileName, VideoId: video.Id}
		db.Postgres.Save(&subtitle)
		subtitle.Save(file)

		//TODO Stop doing this as part of request
		subtitle.ConvertSrtToVtt()
	})

	http.Redirect(response, request, path.ViewVideoPath(video), http.StatusFound)
}
