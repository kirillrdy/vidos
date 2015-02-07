package handler

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func UploadFile(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Redirect(response, request, path.Root, http.StatusFound)
		return
	}

	//TODO fix assumption on buffer size
	request.ParseMultipartForm(1024 * 1024)
	form := request.MultipartForm
	formFiles := form.File[view.FormParamName]

	for _, formFile := range formFiles {
		processFormFile(formFile)
	}

	http.Redirect(response, request, path.Files, http.StatusFound)
}

type uploadedFile struct {
	Filename string
}

func (file uploadedFile) Path() string {
	//TODO use path seperator
	return fmt.Sprintf("%v/%v", fileDir, file.Filename)
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
