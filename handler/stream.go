package handler

import (
	"log"
	"net/http"

	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/path"
)

//TODO serve real filename
//Stream actually streams video content, respecting Range header
func Stream(response http.ResponseWriter, request *http.Request) {

	log.Print(request.Header.Get("Range"))

	video := videoFromRequest(request)

	//response.Header().Set("Content-Type", db.VideoMimeType)
	http.ServeFile(response, request, fs.VideosDataDir+"/"+video.Filepath)
}

func videoFromRequest(request *http.Request) fs.Video {
	return fs.Video{Filepath: request.FormValue(path.ParamKeys.Filepath)}
}
