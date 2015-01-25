package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/db"
)

const Root = "/"
const Videos = "/videos"
const UnencodedVideos = "/unencoded_videos"
const Upload = "/upload"
const Serve = "/serve"
const Download = "/download"
const Reencode = "/reencode"

const NewVideo = "/new_video"

func ServeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Serve, video.Id)
}

func DownloadVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Download, video.Id)
}

func ReencodeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Reencode, video.Id)
}
