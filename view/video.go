package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

const videoThumb css.Class = "video-thumb"

func VideoCss() css.RuleSet {
	return videoThumb.Style(
		//css.Width(size.Px(ffmpeg.ThumbnailWidth)),
		//TODO fix 196, this is youtube's width
		css.Width(size.Px(196)),
	)
}

//For a given video returns its view partial
func Video(video db.Video) html.Node {
	return html.Div().Class(vbox).Children(
		html.H1().Text(video.Filename),
		html.Img().Class(videoThumb).Src(path.ThumbnailPath(video)),
		html.A().Href(path.ServeVideoPath(video)).Text("View"),
	)
}
