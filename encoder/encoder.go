package encoder

import (
	"github.com/kirillrdy/vidos/downloader"
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/util"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var encodingDir = util.VidosDataDirFor("encoding")

func encodeAllfiles() {

	items, err := ioutil.ReadDir(downloader.FilesDir)
	util.LogError(err)

	for _, item := range items {
		if fs.CanBeEncoded(item) {
			encodeFile(item)
		}
	}

}

func encodeFile(file os.FileInfo) {
	filePath := downloader.FilesDir + string(os.PathSeparator) + file.Name()
	encodedName := fs.ChangeExt(file.Name(), fs.Mp4)

	ffmpeg.Encode(filePath, encodingDir+string(os.PathSeparator)+encodedName, func(progress string) {
		log.Printf("Encoding: %s %s\n", file.Name(), progress)
	})

	err := os.Remove(downloader.FilesDir + string(os.PathSeparator) + file.Name())
	util.LogError(err)

	err = os.Rename(encodingDir+string(os.PathSeparator)+encodedName, fs.VideosDataDir+string(os.PathSeparator)+encodedName)
	util.LogError(err)
}

//Start starts a background encoding worker
func Start() {

	go func() {
		for {
			encodeAllfiles()
			//TODO unit all jobs sleep times
			time.Sleep(1 * time.Minute)
		}
	}()
}
