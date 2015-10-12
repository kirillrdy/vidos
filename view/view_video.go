package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/layout"
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

	return html.Div().Class(layout.VBox, layout.Grow, centerItems).Children(
		html.Style().Text(
			css.Stylesheet(
				videoTitle.Style(
					css.FontSize(size.Px(24)),
					css.Height(size.Px(30)),
					css.FlexShrink(0),
				),
				videoPlayer.Style(
					css.MarginBottom(size.Px(10)),
				),
			).String(),
		),
		html.H1().Class(videoTitle).Text(video.Filename),
		html.Video().Class(videoPlayer).Controls().Autoplay().Name("media").Children(videoElementContent...),
		//TODO finish those links, also move them somewhere
		//html.A().Href(path.ManageSubtitlesPath(video)).Text("Download Original"),
		html.A().Href(path.DeleteVideoPath(video)).Text("Delete"),
		html.A().Href(path.ManageSubtitlesPath(video)).Text("Manage Subtitles"),
	)
}
