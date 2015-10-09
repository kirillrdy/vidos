package handler

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

//UploadFile handler accepts file being posted
func UploadFile(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Redirect(response, request, path.Root, http.StatusFound)
		return
	}

	//TODO fix assumption on buffer size
	err := request.ParseMultipartForm(1024 * 1024)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	form := request.MultipartForm
	formFiles := form.File[view.FormParamName]

	for _, formFile := range formFiles {
		processFormFile(formFile)
	}

	http.Redirect(response, request, path.Files.List, http.StatusFound)
}

type uploadedFile struct {
	Filename string
}

func (file uploadedFile) Path() string {
	//TODO use path seperator
	return fmt.Sprintf("%v/%v", downloader.FileDir, file.Filename)
}

func processFormFile(formFile *multipart.FileHeader) {

	//TODO does this needs to be closed ?
	file, err := formFile.Open()
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received %#v", formFile.Filename)

	uploadedFile := uploadedFile{Filename: formFile.Filename}

	//TODO what if Path() didn't exist
	destinationFile, err := os.Create(uploadedFile.Path())
	defer destinationFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(destinationFile, file)
	if n == 0 || err != nil {
		log.Fatal(err)
	}

}
