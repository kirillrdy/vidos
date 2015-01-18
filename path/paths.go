package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/db"
)

const Root = "/"
const Upload = "/upload"
const Serve = "/serve"
const Download = "/download"
const Reencode = "/reencode"

func ServeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Serve, video.Id)
}

func DownloadVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Download, video.Id)
}

func ReencodeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Reencode, video.Id)
}
