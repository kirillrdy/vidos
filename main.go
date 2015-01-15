package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/kirillrdy/nadeshiko/html"
	_ "github.com/lib/pq"
)

func rootHandle(response http.ResponseWriter, request *http.Request) {
	page := html.Html().Children(
		html.Form().Action("/upload").Attribute("enctype", "multipart/form-data").Method("POST").Children(
			html.Input().Type("file").Name("file"),
			html.Input().Type("submit"),
		),
	)
	io.WriteString(response, page.String())
}

func fileUpload(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1)
	form := request.MultipartForm
	formFile := form.File["file"]
	file, err := formFile[0].Open()
	log.Print(formFile[0].Filename)
	if err != nil {
		log.Fatal(err)
	}
	destinationFile, err := os.Create("file")
	io.Copy(destinationFile, file)
	defer file.Close()
	defer destinationFile.Close()
	http.Redirect(response, request, "/", http.StatusFound)
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

	go func() {
		for {
			var stat runtime.MemStats
			runtime.ReadMemStats(&stat)
			log.Print(stat.Alloc / 1024)
			time.Sleep(time.Second)
		}
	}()

	db.AutoMigrate(&Video{})

	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/upload", fileUpload)
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
