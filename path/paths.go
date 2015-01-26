package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/db"
)

const Root = "/"

//Note trailing / is important, due to how "/" http Mux works
const Public = "/public/"
const CssReset = Public + "reset.css"

const Videos = "/videos"
const UnencodedVideos = "/unencoded_videos"
const Upload = "/upload"
const Serve = "/serve"
const Download = "/download"
const Reencode = "/reencode"

const NewVideo = "/new_video"
const Thumbnail = "/thumbnail"

func ServeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Serve, video.Id)
}

func DownloadVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Download, video.Id)
}

func ReencodeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Reencode, video.Id)
}

func ThumbnailPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Thumbnail, video.Id)
}
