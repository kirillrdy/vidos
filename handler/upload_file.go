package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
)

//UploadFile handler accepts file being posted
func UploadFile(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Redirect(response, request, path.Root, http.StatusFound)
		return
	}

	handleMultiFileUpload(response, request, func(file io.ReadCloser, fileName string) {
		defer file.Close()

		uploadedFile := uploadedFile{Filename: fileName}

		//TODO what if Path() didn't exist
		destinationFile, err := os.Create(uploadedFile.Path())
		defer destinationFile.Close()

		//TODO Less Fatal
		if err != nil {
			log.Fatal(err)
		}

		n, err := io.Copy(destinationFile, file)
		if n == 0 || err != nil {
			log.Fatal(err)
		}
	})

	http.Redirect(response, request, path.Files.List, http.StatusFound)
}

type uploadedFile struct {
	Filename string
}

func (file uploadedFile) Path() string {
	//TODO use path seperator
	return fmt.Sprintf("%v/%v", downloader.FilesDir, file.Filename)
}
