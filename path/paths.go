package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/model"
)

const RootPath = "/"
const UploadPath = "/upload"
const ServeFilePath = "/serve"
const DownloadFilePath = "/download"
const ReencodeFilePath = "/reencode"

func ServeVideoPath(video model.Video) string {
	return fmt.Sprintf("%v?id=%v", ServeFilePath, video.Id)
}

func DownloadVideoPath(video model.Video) string {
	return fmt.Sprintf("%v?id=%v", DownloadFilePath, video.Id)
}

func ReencodeVideoPath(video model.Video) string {
	return fmt.Sprintf("%v?id=%v", ReencodeFilePath, video.Id)
}
