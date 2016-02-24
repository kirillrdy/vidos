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
			filePath := downloader.FilesDir + "/" + item.Name()
			encodedName := fs.ChangeExt(item.Name(), fs.Mp4)

			ffmpeg.Encode(filePath, encodingDir+"/"+encodedName, func(progress string) {
				log.Println(progress)
			})

			err := os.Remove(downloader.FilesDir + "/" + item.Name())
			util.LogError(err)

			os.Rename(encodingDir+"/"+encodedName, fs.VideosDataDir+"/"+encodedName)
		}
	}

}

func encodeFile(file os.FileInfo, encodedName string) {

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
