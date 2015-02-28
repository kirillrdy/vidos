package view

import (
	"fmt"
	"log"
	"runtime"
	"syscall"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/display"
	"github.com/sparkymat/webdsl/css/flex"
	"github.com/sparkymat/webdsl/css/overflow"
	"github.com/sparkymat/webdsl/css/size"
)

const AppName = "Видос"
const padding = 10

const siteTitle css.Class = "site-title"
const vbox css.Class = "vbox"
const hbox css.Class = "hbox"
const grow css.Class = "grow"
const noShrink css.Class = "no-shrink"
const wrap css.Class = "wrap"
const linksMenu css.Class = "links-menu"
const centerItems css.Class = "align-items-center"
const headerBar css.Class = "header-bar"
const mainSection css.Class = "main-section"
const statusLine css.Class = "status-line"

func pageStyle() css.CssContainer {
	return css.Stylesheet(
		statusLine.Style(
			css.Height(size.Px(20)),
			css.FlexShrink(0),
			css.PaddingRight(size.Px(padding)),
		),
		siteTitle.Style(
			css.FontSize(size.Px(50)),
		),
		css.AllSelectors(css.Body, css.Html).Style(
			css.Width(size.Percent(100)),
			css.Height(size.Percent(100)),
		),
		vbox.Style(
			css.Display(display.Flex),
			css.FlexDirection(flex.Column),
		),
		hbox.Style(
			css.Display(display.Flex),
			css.FlexDirection(flex.Row),
		),
		wrap.Style(
			css.FlexWrap(flex.Wrap),
		),
		grow.Style(
			css.FlexGrow(1),
		),
		noShrink.Style(
			css.FlexShrink(0),
		),
		centerItems.Style(
			css.AlignItems(css.Center),
		),
		linksMenu.Style(
			css.Width(size.Px(150)),
			css.FlexShrink(0),
			//css.PaddingLeft(size.Px(20)),
		),
		headerBar.Style(
			css.Height(size.Px(70)),
			css.FlexShrink(0),
			css.PaddingLeft(size.Px(padding)),
			css.PaddingRight(size.Px(padding)),
		),
		mainSection.Style(
			css.Overflow(overflow.Auto),
		),
	)
}

func statusLineText() string {

	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	var fsStat syscall.Statfs_t
	err := syscall.Statfs("/", &fsStat)
	if err != nil {
		log.Print(err)
	}

	freeStorage := float64(uint64(fsStat.Bavail)*uint64(fsStat.Bsize)) / float64(1024*1024*1024)
	memoryUsed := float64(memStat.Alloc) / float64(1024*1024)

	return fmt.Sprintf("OS:%v/%v FreeStorage:%.2f Gb, MemUsed: %.2f Mb", runtime.GOOS, runtime.GOARCH, freeStorage, memoryUsed)
}

func Layout(title string, bodyContent ...html.Node) html.Node {
	statusLineText := statusLineText()
	return html.Html().Children(
		html.Head().Children(
			html.Title().Text(title),
			html.Link().Href(path.CSSReset).Rel("stylesheet"),
			html.Style().Text(
				pageStyle().String(),
			),
		),
		html.Body().Class(vbox).Children(
			html.Div().Class(hbox, headerBar, centerItems).Children(
				html.H1().Class(siteTitle).Text(AppName),
				html.Span().Class(grow),
				html.Span().Text(statusLineText),
			),
			html.Div().Class(hbox, grow).Children(
				html.Div().Class(linksMenu, vbox, centerItems).Children(
					html.Div().Class(vbox).Children(
						html.A().Href(path.Videos).Text("Videos"),
						html.A().Href(path.NewVideo).Text("Upload new video"),
						html.A().Href(path.UnencodedVideos).Text("Processing"),
						html.A().Href(path.Files).Text("Files"),
					),
				),
				html.Div().Class(grow, mainSection, vbox).Children(
					bodyContent...,
				),
			),
			html.Div().Class(hbox, statusLine).Children(
				html.Span().Class(grow),
				html.Span().Class().Text(statusLineText),
			),
		),
	)
}
