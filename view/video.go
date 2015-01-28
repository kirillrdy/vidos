package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

const videoThumb css.Class = "video-thumb"
const videoItem css.Class = "video-item"

func VideoCss() css.CssContainer {
	return css.CssContainer{}.Rules(
		videoThumb.Style(
			//css.Width(size.Px(ffmpeg.ThumbnailWidth)),
			//TODO fix 196, this is youtube's width
			css.Width(size.Px(196)),
		),
		videoItem.Style(
			css.Width(size.Px(220)),
		),
	)
}

//For a given video returns its view partial
func Video(video db.Video) html.Node {
	return html.Div().Class(vbox, videoItem).Children(
		html.H1().Text(video.Filename),
		html.Img().Class(videoThumb).Src(path.ThumbnailPath(video)),
		html.A().Href(path.ViewVideoPath(video)).Text("View"),
	)
}
