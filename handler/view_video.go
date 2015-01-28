package handler

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/db"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/view"
)

func ViewVideo(response http.ResponseWriter, request *http.Request) {
	video, err := videoFromRequest(request)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	var subtitles []db.Subtitle
	db.Session.Find(&subtitles, db.Subtitle{VideoId: video.Id})

	var videoElementContent []html.Node

	//TODO dont hardwire type
	videoElementContent = append(videoElementContent, html.Source().Src(path.ServeVideoPath(video)).Type("video/mp4"))

	for _, subtitle := range subtitles {
		track := html.Track().Label("English").Kind("captions").Srclang("en").Src(path.SubtitlePath(subtitle)).Default()
		videoElementContent = append(videoElementContent, track)
	}

	videoElement := html.Video().Controls().Autoplay().Name("media").Children(videoElementContent...)

	inside := html.Div().Children(
		videoElement,
		html.Div().Children(
			html.Span().Text("Upload srt subtitle"),
			html.Form().Action(path.UploadSubtitlePath(video)).Attribute("enctype", "multipart/form-data").Method("POST").Children(
				html.Input().Type("file").Multiple().Name(view.FormParamName),
				html.Input().Type("submit"),
			),
		),
	)
	page := view.Layout(inside)
	io.WriteString(response, page.String())
}
