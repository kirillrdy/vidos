package view

import (
	"github.com/kirillrdy/vidos/ffmpeg"
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

const videoThumb css.Class = "video-thumb"
const videoItem css.Class = "video-item"
const videoLink css.Class = "video-link"

//VideoCSS is css tyles for a single video tile
func VideoCSS() css.CssContainer {
	return css.Stylesheet(
		videoThumb.Style(
			css.Width(size.Px(ffmpeg.ThumbnailWidth/2)),
			//css.Width(size.Percent(30)),
			css.MarginBottom(size.Px(10)),
		),
		videoItem.Style(
			css.Width(size.Px(460)),
			//css.Height(size.Px(170)),
			css.MarginBottom(size.Px(30)),
			css.MarginRight(size.Px(30)),
		),
		videoLink.Style(
			css.WordWrap(css.BreakWord),
		),
	)
}

//Video For a given video returns its view partial
func Video(video fs.Video) html.Node {
	return html.Div().Class(flex.VBox, videoItem, flex.CenterItems).Children(
		//html.Img().Class(videoThumb).Src(path.ThumbnailPath(video)),
		html.A().Class(videoLink).Href(path.ViewVideoPath(video)).Text(video.Filepath),
	)
}
