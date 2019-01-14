package fs

import (
	"github.com/kirillrdy/vidos/util"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path/filepath"
)

//VideosDataDir This is where all servable vidoes are stored, including subtitles
var VideosDataDir = util.VidosDataDirFor("videos")

//Video represents a video file
type Video struct {
	Filepath string
}

//MimeType returns mimetype for a video
func (video Video) MimeType() string {
	ext := filepath.Ext(video.Filepath)
	// this is because golang's built in mime database only has very limited list of mime types
	if ext == ".mp4" {
		return "video/mp4"
	}
	mime := mime.TypeByExtension(ext)
	if mime == "" {
		log.Println("WARNING: failed to detect mime type")
	}

	return mime
}

//Delete deletes video from fs, also will delete any related metadata
//eg subtitles
func (video Video) Delete() error {
	return os.Remove(video.Filepath)
}

//Filename returns filename of a file
func (video Video) Filename() string {
	return filepath.Base(video.Filepath)
}

//Videos returs all streamable videos in the VideosDataDir
func Videos() ([]Video, error) {
	var videos []Video
	files, err := ioutil.ReadDir(VideosDataDir)
	if err != nil {
		return videos, err
	}

	for _, file := range files {
		if canBeStreamed(file) {
			videos = append(videos, Video{Filepath: file.Name()})
		}
	}
	return videos, nil
}

// Save reads content from the reader and writes it to filepath of video
func (video Video) Save(reader io.ReadCloser) error {

	destinationFile, err := os.Create(video.Filepath)

	if err != nil {
		return err
	}

	n, err := io.Copy(destinationFile, reader)
	if n == 0 || err != nil {
		return err
	}

	err = reader.Close()
	if err != nil {
		return err
	}
	err = destinationFile.Close()
	if err != nil {
		return err
	}
	return nil
}
