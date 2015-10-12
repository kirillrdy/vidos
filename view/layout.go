package view

import (
	"fmt"
	"log"
	"runtime"
	"syscall"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/layout"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/overflow"
	"github.com/sparkymat/webdsl/css/size"
)

const appName = "Видос"
const padding = 10

const siteTitle css.Class = "site-title"
const linksMenu css.Class = "links-menu"
const centerItems css.Class = "align-items-center"
const headerBar css.Class = "header-bar"
const mainSection css.Class = "main-section"
const statusLine css.Class = "status-line"
const menuItem css.Class = "menu-item"

func pageStyle() css.CssContainer {
	return css.Stylesheet(
		statusLine.Style(
			css.Height(size.Px(30)),
			css.FlexShrink(0),
			css.PaddingRight(size.Px(padding)),
		),
		siteTitle.Style(
			css.FontSize(size.Px(50)),
		),
		centerItems.Style(
			css.AlignItems(css.Center),
		),
		linksMenu.Style(
			css.Width(size.Px(197)),
			css.FlexShrink(0),
		),
		headerBar.Style(
			css.Height(size.Px(100)),
			css.FlexShrink(0),
		),
		mainSection.Style(
			css.Overflow(overflow.Auto),
			css.Padding(size.Px(10)),
		),
		menuItem.Style(
			css.MarginTop(size.Px(padding)),
			css.MarginBottom(size.Px(padding)),
		),
	)
}

func statusLineText() string {

	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	var fsStat syscall.Statfs_t
	//TODO dont use '/', but use datadir
	err := syscall.Statfs("/", &fsStat)
	if err != nil {
		log.Print(err)
	}

	freeStorage := float64(uint64(fsStat.Bavail)*uint64(fsStat.Bsize)) / float64(1024*1024*1024)
	memoryUsed := float64(memStat.Alloc) / float64(1024*1024)

	return fmt.Sprintf("OS:%v/%v FreeStorage:%.2f Gb, MemUsed: %.2f Mb", runtime.GOOS, runtime.GOARCH, freeStorage, memoryUsed)
}

//Layout returns the main layout of the application
//TODO fix regression with overflow in the main grow section
func Layout(title string, bodyContent ...html.Node) html.Node {

	title = fmt.Sprintf("%v - %v", appName, title)

	statusLineText := statusLineText()
	return html.Html().Children(
		html.Head().Children(
			html.Title().Text(title),
			html.Link().Href(path.CSSReset).Rel("stylesheet"),
			html.Style().Text(
				pageStyle().String(),
			),
			layout.Styles(),
		),
		html.Body().Class(layout.VBox).Children(
			html.Div().Class(layout.HBox, headerBar, centerItems).Children(
				html.Span().Class(layout.Grow),
				html.H1().Class(siteTitle).Text(appName),
				html.Span().Class(layout.Grow),
			),
			html.Div().Class(layout.HBox, layout.Grow).Children(
				html.Div().Class(linksMenu, layout.VBox, centerItems).Children(
					html.Div().Class(layout.VBox).Children(
						html.A().Class(menuItem).Href(path.Videos.List).Text("Videos"),
						html.A().Class(menuItem).Href(path.Videos.New).Text("Upload new video"),
						html.A().Class(menuItem).Href(path.Videos.Unencoded).Text("Processing"),
						html.A().Class(menuItem).Href(path.Files.List).Text("Files"),
						html.A().Class(menuItem).Href(path.Torrents).Text("Torrents"),
						html.A().Class(menuItem).Href(path.AddMagnetLink).Text("Add Magnet link"),
					),
				),
				html.Div().Class(layout.Grow, mainSection, layout.VBox).Children(
					bodyContent...,
				),
			),
			html.Div().Class(layout.HBox, statusLine).Children(
				html.Span().Class(layout.Grow),
				html.Span().Class().Text(statusLineText),
			),
		),
	)
}
