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

const siteTitle css.Class = "site-title"
const vbox css.Class = "vbox"
const hbox css.Class = "hbox"
const grow css.Class = "grow"
const wrap css.Class = "wrap"
const linksMenu css.Class = "links-menu"
const centerItems css.Class = "align-items-center"
const headerBar css.Class = "header-bar"
const mainSection css.Class = "main-section"

func pageStyle() css.CssContainer {
	return css.Stylesheet(
		siteTitle.Style(
			css.FontSize(size.Px(50)),
		),
		css.AllSelectors(css.Body, css.Html).Style(
			css.Width(size.Percent(100)),
			css.Height(size.Percent(100)),
		),
		css.ChildSelector(vbox, css.All).Style(
			css.FlexShrink(0),
		),
		css.ChildSelector(hbox, css.All).Style(
			css.FlexShrink(0),
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
		centerItems.Style(
			css.AlignItems(css.Center),
		),
		linksMenu.Style(
			css.Width(size.Px(180)),
		),
		headerBar.Style(
			css.Height(size.Px(70)),
			css.PaddingLeft(size.Px(10)),
			css.PaddingRight(size.Px(10)),
		),
		mainSection.Style(
			css.Overflow(overflow.Auto),
			css.FlexShrink(1),
		),
	)
}

func statusLine() string {

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
	return html.Html().Children(
		html.Head().Children(
			html.Title().Text(title),
			html.Link().Href(path.CssReset).Rel("stylesheet"),
			html.Style().Text(
				pageStyle().String(),
			),
		),
		html.Body().Class(vbox).Children(
			html.Div().Class(hbox, headerBar, centerItems).Children(
				html.H1().Class(siteTitle).Text(AppName),
				html.Span().Class(grow),
				html.Span().Text(statusLine()),
			),
			html.Div().Class(hbox, grow).Children(
				html.Div().Class(linksMenu, vbox, centerItems).Children(
					html.A().Href(path.Videos).Text("Videos"),
					html.A().Href(path.NewVideo).Text("Upload new video"),
					html.A().Href(path.UnencodedVideos).Text("Processing"),
					html.A().Href(path.Files).Text("Files"),
				),
				html.Div().Class(grow, mainSection, vbox).Children(
					bodyContent...,
				),
			),
		),
	)
}
