package videos

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

//This is where all servable vidoes are stored, including subtitles
var VideosDataDir string

func All() ([]Video, error) {
	var videos []Video
	files, err := ioutil.ReadDir(VideosDataDir)
	if err != nil {
		return videos, err
	}

	for _, file := range files {
		videos = append(videos, Video{Filename: file.Name()})

	}
	return videos, nil
}

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

func init() {

	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}

	VideosDataDir = user.HomeDir + "/./vidos/videos"
	err = os.MkdirAll(VideosDataDir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
}
