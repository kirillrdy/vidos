package view

import (
	"fmt"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

func ViewVideo(video db.Video, subtitles []db.Subtitle) html.Node {
	var videoElementContent []html.Node

	videoElementContent = append(videoElementContent, html.Source().Src(path.ServeVideoPath(video)).Type(db.VideoMimeType))

	for _, subtitle := range subtitles {
		//TODO fix hardwired language
		track := html.Track().Label("English").Kind("captions").Srclang("en").Src(path.SubtitlePath(subtitle)).Default()
		videoElementContent = append(videoElementContent, track)
	}

	var videoPlayer css.Class = "video-player"
	var videoTitle css.Class = "video-title"

	inside := html.Div().Class(vbox, grow, centerItems).Children(
		html.Style().Text(
			css.Stylesheet(
				videoTitle.Style(
					css.FontSize(size.Px(24)),
					css.Height(size.Px(30)),
					css.FlexShrink(0),
				),
				videoPlayer.Style()).String(),
		),
		html.H1().Class(videoTitle).Text(video.Filename),
		html.Video().Class(videoPlayer).Controls().Autoplay().Name("media").Children(videoElementContent...),
		//TODO finish those links, also move them somewhere
		//html.A().Href(path.ManageSubtitlesPath(video)).Text("Download Original"),
		//html.A().Href(path.ManageSubtitlesPath(video)).Text("Delete"),
		html.A().Href(path.ManageSubtitlesPath(video)).Text("Manage Subtitles"),
	)

	//Title is important here for chromecast support :-)
	title := fmt.Sprintf("%v - %v", AppName, video.Filename)
	return Layout(title, inside)
}
