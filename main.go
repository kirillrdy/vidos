package main

import (
	"fmt"
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

	var trs []html.Node
	var videos []Video
	result := db.Find(&videos)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	for _, video := range videos {
		tr := html.Tr().Children(
			html.Td().Text(video.IdString()),
			html.Td().Text(video.Filename),
		)
		trs = append(trs, tr)
	}

	page := html.Html().Children(
		html.Table().Children(
			html.Thead().Children(
				html.Tr().Children(
					html.Th().Text("Id"),
					html.Th().Text("File name"),
				),
			),

			html.Tbody().Children(
				trs...,
			),
		),
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

	video := Video{Filename: formFile[0].Filename}
	db.Save(&video)
	video.mkdir()

	log.Print(formFile[0].Filename)

	destinationFile, err := os.Create(video.filePath())

	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(destinationFile, file)
	if n == 0 || err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer destinationFile.Close()
	http.Redirect(response, request, "/", http.StatusFound)
}

func startMemoryMonitoring() {
	go func() {
		for {
			var stat runtime.MemStats
			runtime.ReadMemStats(&stat)
			log.Printf("%vKb", stat.Alloc/1024)
			time.Sleep(time.Second)
		}
	}()
}

type Video struct {
	Id       uint64
	Filename string
}

const dataDir = "data"

func (video Video) dirPath() string {
	return fmt.Sprintf("%v/%v", dataDir, video.Id)
}
func (video Video) filePath() string {
	return fmt.Sprintf("%v/%v", video.dirPath(), video.Filename)
}

func (video Video) IdString() string {
	return fmt.Sprint(video.Id)
}

func (video Video) mkdir() {
	err := os.MkdirAll(video.dirPath(), os.ModePerm)
	if err != nil {
		log.Print(err)
	}
}

var db gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "dbname=vidos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Video{})

	//startMemoryMonitoring()

	http.HandleFunc("/", rootHandle)
	http.HandleFunc(uploadPath, fileUpload)
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
