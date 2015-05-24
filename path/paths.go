package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/db"
)

// Root page, which actually redirects to a more useful page
const Root = "/"

//Note trailing / is important, due to how "/" http Mux works
const Public = "/public/"
const CSSReset = Public + "reset.css"

const Videos = "/videos"
const ViewVideo = "/view_video"
const UnencodedVideos = "/unencoded_videos"
const Upload = "/upload"
const UploadSubtitle = "/upload_subtitle"
const Serve = "/serve"
const Download = "/download"
const Reencode = "/reencode"
const Delete = "/delete"

const NewVideo = "/new_video"
const Thumbnail = "/thumbnail"
const Subtitle = "/subtitle.vtt"

const Files = "/files"
const UploadFile = "/upload_file"

const ManageSubtitles = "/subtitles"

func ServeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Serve, video.Id)
}

func DownloadVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Download, video.Id)
}

func ReencodeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Reencode, video.Id)
}

func DeleteVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Delete, video.Id)
}

func ThumbnailPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", Thumbnail, video.Id)
}

func ViewVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", ViewVideo, video.Id)
}

func UploadSubtitlePath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", UploadSubtitle, video.Id)
}

func SubtitlePath(subtitle db.Subtitle) string {
	return fmt.Sprintf("%v?id=%v", Subtitle, subtitle.Id)
}

func ManageSubtitlesPath(video db.Video) string {
	return fmt.Sprintf("%v?id=%v", ManageSubtitles, video.Id)
}
