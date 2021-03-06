package view

import (
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/fs"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/web/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

//VideoShowPage view renders View Video View :-)
func VideoShowPage(video fs.Video) html.Node {
	var videoElementContent []html.Node

	videoElementContent = append(videoElementContent, html.Source().Src(path.StreamVideoPath(video)).Type(video.MimeType()))

	// for _, subtitle := range subtitles {
	// 	//TODO fix hardwired language
	// 	track := html.Track().Label("English").Kind("captions").Srclang("en").Src(path.SubtitlePath(subtitle)).Default()
	// 	videoElementContent = append(videoElementContent, track)
	// }

	var videoPlayer css.Class = "video-player"
	var videoTitle css.Class = "video-title"

	page := html.Div().Class(flex.VBox, flex.Grow, flex.CenterItems).Children(
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
		html.H1().Class(videoTitle).Text(video.Filepath),
		html.Video().Class(videoPlayer).Controls().Autoplay().Name("media").Children(videoElementContent...),
	//TODO finish those links, also move them somewhere
	//html.A().Href(path.ManageSubtitlesPath(video)).Text("Download Original"),
	// html.A().Href(path.DeleteVideoPath(video)).Text("Delete"),
	// 	html.A().Href(path.ManageSubtitlesPath(video)).Text("Manage Subtitles"),
	)

	return application.NewPage(video.Filename(), path.ViewVideoPath(video)).ToHTML(page)
}
