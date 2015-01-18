package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/model"
)

const Root = "/"
const Upload = "/upload"
const Serve = "/serve"
const Download = "/download"
const Reencode = "/reencode"

func ServeVideoPath(video model.Video) string {
	return fmt.Sprintf("%v?id=%v", Serve, video.Id)
}

func DownloadVideoPath(video model.Video) string {
	return fmt.Sprintf("%v?id=%v", Download, video.Id)
}

func ReencodeVideoPath(video model.Video) string {
	return fmt.Sprintf("%v?id=%v", Reencode, video.Id)
}
