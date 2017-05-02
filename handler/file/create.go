package file

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

//Create handler accepts file being posted
func Create(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Redirect(response, request, path.Root, http.StatusFound)
		return
	}

	err := handleMultiFileUpload(response, request, func(file io.ReadCloser, fileName string) error {
		defer file.Close()

		uploadedFile := uploadedFile{Filename: fileName}

		destinationFile, err := os.Create(uploadedFile.Path())

		if err != nil {
			return err
		}
		defer destinationFile.Close()

		n, err := io.Copy(destinationFile, file)
		if n == 0 || err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO perhaps have own way of redirecting that uses web.Path type
	http.Redirect(response, request, path.Files.List.String(), http.StatusFound)
}

type uploadedFile struct {
	Filename string
}

func (file uploadedFile) Path() string {
	return downloader.FilesDir + string(os.PathSeparator) + file.Filename
}

func handleMultiFileUpload(response http.ResponseWriter, request *http.Request, fileProcessor func(io.ReadCloser, string) error) error {

	//TODO too much duplication
	//TODO fix assumption on buffer size
	err := request.ParseMultipartForm(1024 * 1024)
	if err != nil {
		return err
	}
	form := request.MultipartForm
	formFiles := form.File[view.FileParamName]

	for _, formFile := range formFiles {
		//TODO does this needs to be closed ?
		file, err := formFile.Open()

		//TODO don't Fatal
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received %#v", formFile.Filename)
		err = fileProcessor(file, formFile.Filename)
		if err != nil {
			return err
		}
	}
	return nil
}
