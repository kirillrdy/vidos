package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kirillrdy/nadeshiko/html"
	_ "github.com/lib/pq"
)

const formParamName = "file"
const uploadPath = "/upload"

func rootHandle(response http.ResponseWriter, request *http.Request) {
	page := html.Html().Children(
		html.Form().Action(uploadPath).Attribute("enctype", "multipart/form-data").Method("POST").Children(
			html.Input().Type("file").Name(formParamName),
			html.Input().Type("submit"),
		),
	)
	io.WriteString(response, page.String())
}

func fileUpload(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1024 * 1024)
	form := request.MultipartForm
	formFile := form.File[formParamName]
	//TODO handle no file submitted and only 1 file submitted
	file, err := formFile[0].Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(formFile[0].Filename)

	destinationFile, err := os.Create("file")
	io.Copy(destinationFile, file)
	defer file.Close()
	defer destinationFile.Close()
	http.Redirect(response, request, "/", http.StatusFound)
}

func startMemoryMonitoring() {
	go func() {
		for {
			var stat runtime.MemStats
			runtime.ReadMemStats(&stat)
			log.Print(stat.Alloc / 1024)
			time.Sleep(time.Second)
		}
	}()
}

type Video struct {
	Id       uint64
	filename string
}

var db gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "dbname=vidos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Video{})

	startMemoryMonitoring()

	http.HandleFunc("/", rootHandle)
	http.HandleFunc(uploadPath, fileUpload)
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
