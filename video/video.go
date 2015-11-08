package video

import (
	"github.com/kirillrdy/vidos/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Video represents a video file
type Video struct {
	Filename string
}

//VideosDataDir This is where all servable vidoes are stored, including subtitles
var VideosDataDir = util.VidosDataDirFor("videos")

//All returs all streamable videos in the VideosDataDir
func All() ([]Video, error) {
	var videos []Video
	files, err := ioutil.ReadDir(VideosDataDir)
	if err != nil {
		return videos, err
	}

	for _, file := range files {
		if canBeStreamed(file) {
			videos = append(videos, Video{Filename: file.Name()})
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
