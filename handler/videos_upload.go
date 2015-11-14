package handler

// import (
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/kirillrdy/vidos/db"
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

// func handleMultiFileUpload(response http.ResponseWriter, request *http.Request, fileProcessor func(io.ReadCloser, string) error) error {

// 	//TODO too much duplication
// 	//TODO fix assumption on buffer size
// 	err := request.ParseMultipartForm(1024 * 1024)
// 	if err != nil {
// 		return err
// 	}
// 	form := request.MultipartForm
// 	formFiles := form.File[view.FileParamName]

// 	for _, formFile := range formFiles {
// 		//TODO does this needs to be closed ?
// 		file, err := formFile.Open()

// 		//TODO don't Fatal
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("Received %#v", formFile.Filename)
// 		err = fileProcessor(file, formFile.Filename)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func processVideoFromFile(file io.ReadCloser, filename string) error {

// 	video := db.Video{Filename: filename}
// 	db.Postgres.Save(&video)
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

// 	//This can block so do in goroutine
// 	//TODO potentially dangerous
// 	go func() {
// 		db.EncodeVideo <- video.Id
// 	}()

// 	return nil
// }
