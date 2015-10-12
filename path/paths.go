package path

import (
	"fmt"

	"github.com/kirillrdy/vidos/db"
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
	New       string
	Create    string
	Show      string
	Stream    string
	Download  string
	Delete    string
	Thumbnail string
}{
	"/videos/list",
	"/videos/unencoded",
	"/videos/new",
	"/videos/create",
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
const AddFileForEncoding = "/add_file_for_encoding"

var ParamKeys = struct {
	ID       string
	Filepath string
	Path     string
}{
	"id",
	"filepath",
	"path",
}

func DeleteFileOrDirectoryPath(filename string) string {
	return fmt.Sprintf("%v?%v=%v", DeleteFileOrDirectory, ParamKeys.Filepath, url.QueryEscape(filename))
}

func AddFileForEncodingPath(filename string) string {
	return fmt.Sprintf("%v?%v=%v", AddFileForEncoding, ParamKeys.Filepath, url.QueryEscape(filename))
}

func ViewFilesPath(dirName string) string {
	return fmt.Sprintf("%v?%v=%v", Files.List, ParamKeys.Path, url.QueryEscape(dirName))
}

func ServeVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Stream, ParamKeys.ID, video.Id)
}

func DownloadVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Download, ParamKeys.ID, video.Id)
}

func DeleteVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Delete, ParamKeys.ID, video.Id)
}

func ThumbnailPath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Thumbnail, ParamKeys.ID, video.Id)
}

func ViewVideoPath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", Videos.Show, ParamKeys.ID, video.Id)
}

func UploadSubtitlePath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", UploadSubtitle, ParamKeys.ID, video.Id)
}

func SubtitlePath(subtitle db.Subtitle) string {
	return fmt.Sprintf("%v?%v=%v", Subtitle, ParamKeys.ID, subtitle.Id)
}

func ManageSubtitlesPath(video db.Video) string {
	return fmt.Sprintf("%v?%v=%v", ManageSubtitles, ParamKeys.ID, video.Id)
}
