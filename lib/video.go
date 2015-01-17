package lib

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kirillrdy/vidos/ffmpeg"
)

const dataDir = "data"

type Video struct {
	Id       uint64
	Filename string
	Encoded  bool
}

func (video Video) dirPath() string {
	return fmt.Sprintf("%v/%v", dataDir, video.Id)
}
func (video Video) filePath() string {
	return fmt.Sprintf("%v/%v", video.dirPath(), video.Filename)
}

func (video Video) IdString() string {
	return fmt.Sprint(video.Id)
}

func (video Video) Mkdir() {
	err := os.MkdirAll(video.dirPath(), os.ModePerm)
	if err != nil {
		log.Print(err)
	}
}

func (video Video) StartEncoding() {
	go func() {
		video.Encode()
		video.Encoded = true
		Db.Save(&video)
	}()
}

func (video Video) Encode() {
	ffmpeg.Encode(video.filePath(), video.encodedPath())
}

func (video Video) Duration() string {
	return ffmpeg.Duration(video.filePath())
}

func (video Video) encodedPath() string {
	return video.filePath() + ".mp4"
}

//TODO get rid of log.Fatal
func (video Video) Save(reader io.ReadCloser) {

	video.Mkdir()

	destinationFile, err := os.Create(video.filePath())

	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(destinationFile, reader)
	if n == 0 || err != nil {
		log.Fatal(err)
	}

	defer reader.Close()
	defer destinationFile.Close()
}
