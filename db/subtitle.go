package db

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kirillrdy/vidos/srt2webvtt"
)

type Subtitle struct {
	Id       uint64
	VideoId  uint64
	Filename string
}

func (subtitle Subtitle) dataDirPath() string {
	return fmt.Sprintf("%v/subtitles/%v", dataDir, subtitle.Id)
}
func (subtitle Subtitle) FilePath() string {
	return fmt.Sprintf("%v/%v", subtitle.dataDirPath(), subtitle.Filename)
}

func (subtitle Subtitle) vttFilename() string {
	return subtitle.Filename + ".vtt"
}

func (subtitle Subtitle) VttFilePath() string {
	return fmt.Sprintf("%v/%v", subtitle.dataDirPath(), subtitle.vttFilename())
}

func (subtitle Subtitle) ConvertSrtToVtt() {
	srt2webvtt.Convert(subtitle.FilePath(), subtitle.VttFilePath())
}
func (subtitle Subtitle) mkdir() {
	err := os.MkdirAll(subtitle.dataDirPath(), os.ModePerm)
	if err != nil {
		log.Print(err)
	}
}

//TODO rewrite this to use attachment ( or file upload whatever its called ) interface
func (subtitle Subtitle) Delete() error {
	err := os.Remove(subtitle.FilePath())
	if err != nil {
		return err
	}
	err = os.Remove(subtitle.VttFilePath())
	if err != nil {
		return err
	}
	result := Session.Delete(&subtitle)
	return result.Error
}

//TODO get rid of log.Fatal
//TODO this is very similar to method in video.go
func (subtitle Subtitle) Save(reader io.ReadCloser) {

	subtitle.mkdir()

	destinationFile, err := os.Create(subtitle.FilePath())

	if err != nil {
		log.Println("Subtitle/Save()")
		log.Fatal(err)
	}

	n, err := io.Copy(destinationFile, reader)
	if n == 0 || err != nil {
		log.Fatal(err)
	}

	defer reader.Close()
	defer destinationFile.Close()
}
