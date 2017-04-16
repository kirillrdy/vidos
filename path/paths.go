package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/fs"
	"net/url"
)

// Root page, which actually redirects to a more useful page
const Root = "/"

//Note trailing / is important, due to how "/" http Mux works
const Public = "/public/"
const CSSReset = Public + "reset.css"

var Videos = struct {
	List      string
	Unencoded string
	Show      string
	Stream    string
	Download  string
	Delete    string
	Thumbnail string
}{
	"/videos/list",
	"/videos/unencoded",
	"/videos/show",
	"/videos/stream",
	"/videos/download",
	"/videos/delete",
	"/videos/thumbnail",
}

var Files = struct {
	List   string
	Upload string
}{
	"/files/list",
	"/files/upload",
}

const UploadSubtitle = "/upload_subtitle"

const Subtitle = "/subtitle.vtt"

//UploadFile is path for handling of the manual file upload, not to be confused with Upload
const UploadFile = "/upload_file"
const DeleteFileOrDirectory = "/delete_file"

const ManageSubtitles = "/subtitles"
const Torrents = "/torrents"
const TorrentStatus = "/torrent_status"
const AddMagnetLink = "/add_magnet_link"

var ParamKeys = struct {
	ID       string
	Filepath string
	Path     string
}{
	"id",
	"filepath",
	"path",
}

//TODO refactor all those to use url.Values
func DeleteFileOrDirectoryPath(filename string) string {
	return fmt.Sprintf("%v?%v=%v", DeleteFileOrDirectory, ParamKeys.Filepath, url.QueryEscape(filename))
}

func ViewFilesPath(dirName string) string {
	return fmt.Sprintf("%v?%v=%v", Files.List, ParamKeys.Path, url.QueryEscape(dirName))
}

//StreamVideoPath is path where actual video being streamed from
func StreamVideoPath(video fs.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Stream, ParamKeys.Filepath, url.QueryEscape(video.Filepath))
}

// func DownloadVideoPath(video db.Video) string {
// 	return fmt.Sprintf("%v?%v=%v", Videos.Download, ParamKeys.ID, video.Id)
// }

// func DeleteVideoPath(video db.Video) string {
// 	return fmt.Sprintf("%v?%v=%v", Videos.Delete, ParamKeys.ID, video.Id)
// }

// func ThumbnailPath(video video.Video) string {
// 	return fmt.Sprintf("%v?%v=%v", Videos.Thumbnail, ParamKeys.ID, video.Filename)
// }

//ViewFilesPath is path where video player is rendered
func ViewVideoPath(video fs.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Show, ParamKeys.Filepath, url.QueryEscape(video.Filepath))
}

// func UploadSubtitlePath(video db.Video) string {
// 	return fmt.Sprintf("%v?%v=%v", UploadSubtitle, ParamKeys.ID, video.Id)
// }

// func SubtitlePath(subtitle db.Subtitle) string {
// 	return fmt.Sprintf("%v?%v=%v", Subtitle, ParamKeys.ID, subtitle.Id)
// }

// func ManageSubtitlesPath(video db.Video) string {
// 	return fmt.Sprintf("%v?%v=%v", ManageSubtitles, ParamKeys.ID, video.Id)
// }
