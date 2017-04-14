package handler

// import (
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/kirillrdy/vidos/path"
// 	"github.com/kirillrdy/vidos/view"
// )

// //VideosUpload is a handler that deals with
// // recieving videos via http post and converting them and adding to list of videos
// func VideosUpload(response http.ResponseWriter, request *http.Request) {

// 	if request.Method != "POST" {
// 		http.Redirect(response, request, path.Root, http.StatusFound)
// 		return
// 	}

// 	err := handleMultiFileUpload(response, request, processVideoFromFile)
// 	if err != nil {
// 		http.Error(response, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(response, request, path.Root, http.StatusFound)
// }

// func processVideoFromFile(file io.ReadCloser, filename string) error {

// 	video.Save(file)

// 	//TODO Stop doing this as part of request
// 	err := video.CalculateDuration()
// 	if err != nil {
// 		return err
// 	}
// 	err = video.GenerateThumbnail()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
