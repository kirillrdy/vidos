package fs

import (
	"github.com/kirillrdy/vidos/util"
	"io/ioutil"
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
	return mime.TypeByExtension(ext)
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

func canBeStreamed(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}
	ext := filepath.Ext(file.Name())
	return ext == ".mp4"
}

//CanBeEncoded for a given os.FileInfo returns if the file can be encoded using ffmpeg
func CanBeEncoded(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}
	ext := filepath.Ext(file.Name())
	if ext == ".mp4" || ext == ".avi" || ext == ".mkv" {
		return true
	}
	return false
}
