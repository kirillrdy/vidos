package db

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kirillrdy/vidos/ffmpeg"
)

const dataDir = "data"
const thumbnailFilename = "thumb.png"

//VideoMimeType returns a mimetype for a mp4 video
const VideoMimeType = "video/mp4"

type Video struct {
	Id       uint64
	Filename string
	Encoded  bool
	Progress string
	Duration string
}

func (video Video) dataDirPath() string {
	return fmt.Sprintf("%v/videos/%v", dataDir, video.Id)
}

func (video Video) filePath() string {
	return fmt.Sprintf("%v/%v", video.dataDirPath(), video.Filename)
}

func (video Video) IDString() string {
	return fmt.Sprint(video.Id)
}

func (video Video) mkdir() {
	err := os.MkdirAll(video.dataDirPath(), os.ModePerm)
	if err != nil {
		log.Print(err)
	}
}

func (video Video) Encode() {
	update := func(timeProgress string) {
		video.Progress = timeProgress
		//TODO errors
		Postgres.Save(&video)
	}
	ffmpeg.Encode(video.filePath(), video.EncodedPath(), update)

	video.Encoded = true
	video.Progress = ""
	//TODO errors
	Postgres.Save(&video)

	//remove original
	err := os.Remove(video.filePath())
	if err != nil {
		log.Print("video/Encode() couldn't remove original file")
		log.Panic(err)
	}
}

//CalculateDuration calculates and sotores the duration of the video
func (video *Video) CalculateDuration() error {
	var err error
	video.Duration, err = ffmpeg.Duration(video.filePath())
	if err != nil {
		return err
	}

	//TODO handle erorors on postgres.save
	Postgres.Save(video)
	return nil
}

//GenerateThumbnail generates a thumbnamil image for a video
func (video *Video) GenerateThumbnail() error {
	return ffmpeg.Thumbnail(video.filePath(), video.ThumbnailPath())
}

func (video Video) EncodedPath() string {
	//TODO do basename to strip out ext
	//TODO somehow note the importance of the mp4 here,
	// since its how ffmpeg decides to encode to mp4
	return video.filePath() + ".mp4"
}

func (video Video) ThumbnailPath() string {
	return fmt.Sprintf("%v/%v", video.dataDirPath(), thumbnailFilename)
}

//TODO get rid of log.Fatal
func (video Video) Save(reader io.ReadCloser) {

	video.mkdir()

	destinationFile, err := os.Create(video.filePath())

	if err != nil {
		log.Println("Video/Save()")
		log.Fatal(err)
	}

	n, err := io.Copy(destinationFile, reader)
	if n == 0 || err != nil {
		log.Println("Video/Save()/io.Copy")
		log.Fatal(err)
	}

	defer reader.Close()
	defer destinationFile.Close()
}

func (video Video) Delete() {
	//TODO errors
	Postgres.Delete(&video)
	err := os.RemoveAll(video.dataDirPath())
	if err != nil {
		log.Print(err)
	}
}

var EncodeVideo = make(chan uint64, 100) //TODO why 100

func QueueAllUnEncodedVideos() {

	var videos []Video

	//TODO "id" should be a constant
	Postgres.Order("id").Find(&videos)

	for _, video := range videos {
		if video.Encoded == false {
			EncodeVideo <- video.Id
			log.Printf("Added video for encoding to queue, %v", video.Filename)
		}
	}
}

func init() {
	go func() {
		for {
			id := <-EncodeVideo

			var video Video
			Postgres.Find(&video, id)
			video.Encode()

		}
	}()
}
