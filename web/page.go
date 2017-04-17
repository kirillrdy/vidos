package web

import (
	"fmt"
	"runtime"
	"syscall"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/util"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/overflow"
	"github.com/sparkymat/webdsl/css/size"
	"github.com/sparkymat/webdsl/css/textdecoration"
)

var padding = size.Px(10)

const siteTitle css.Class = "site-title"
const linksMenu css.Class = "links-menu"
const headerBar css.Class = "header-bar"
const mainSection css.Class = "main-section"
const statusLine css.Class = "status-line"
const menuItem css.Class = "menu-item"
const selectedMenuItem css.Class = "selected-menu-item"

func pageStyle() css.CssContainer {
	return css.Stylesheet(
		css.Body.Style(
			css.FontFamily("'Helvetica Neue',Helvetica,Arial,sans-serif;"),
			css.Color(color.ColorRGB{Red: 33, Green: 33, Blue: 33}),
			css.BackgroundColor(color.ColorRGB{Red: 242, Green: 242, Blue: 242}),
		),
		css.Element("table").Style(css.Width(size.Percent(100))),
		statusLine.Style(
			css.Height(size.Px(30)),
			css.FlexShrink(0),
			css.PaddingRight(padding),
		),
		siteTitle.Style(
			css.FontSize(size.Px(25)),
			css.Color(color.White),
			css.PaddingLeft(padding),
		),
		linksMenu.Style(
			css.Width(size.Px(197)),
			css.FlexShrink(0),
		),
		headerBar.Style(
			css.Height(size.Px(56)),
			css.BackgroundColor(color.ColorRGB{Red: 66, Green: 133, Blue: 244}),
			css.FlexShrink(0),
			css.BoxShadow(size.Px(0), size.Px(4), size.Px(4), color.Gray),
		),
		mainSection.Style(
			css.Overflow(overflow.Auto),
			css.Padding(padding),
		),
		menuItem.Style(
			css.Padding(size.Px(15)),
			css.TextDecoration(textdecoration.None),
			css.Color(color.ColorRGB{Red: 97, Green: 97, Blue: 97}),
		),
		css.SelectorWithPseudoClass{Element: menuItem, PseudoClass: css.Hover}.Style(
			css.BackgroundColor(color.ColorRGBA{Red: 0, Green: 0, Blue: 0, Alpha: 0.05}),
		),
		selectedMenuItem.Style(
			css.BackgroundColor(color.ColorRGBA{Red: 0, Green: 0, Blue: 0, Alpha: 0.05}),
			css.FontWeightBold(),
			css.Color(color.ColorRGB{Red: 33, Green: 33, Blue: 33}),
		),
	)
}

func statusLineText() string {

	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	var fsStat syscall.Statfs_t
	err := syscall.Statfs(util.VidosDataDir, &fsStat)
	util.LogError(err)

	// note that forceful type casting necessary here due to different types on different platforms
	freeStorage := float64(uint64(fsStat.Bavail)*uint64(fsStat.Bsize)) / float64(1024*1024*1024)
	memoryUsed := float64(memStat.Alloc) / float64(1024*1024)

	return fmt.Sprintf("OS:%v/%v FreeStorage:%.2f Gb, MemUsed: %.2f Mb", runtime.GOOS, runtime.GOARCH, freeStorage, memoryUsed)
}

//Page returns the main layout of the application
//TODO fix regression with overflow in the main grow section
func Page(title string, subTitle string, bodyContent ...html.Node) html.Node {

	if subTitle != "" {
		title = fmt.Sprintf("%v - %v", title, subTitle)
	}

	statusLineText := statusLineText()
	return html.Html().Children(
		html.Head().Children(
			html.Title().Text(title),
			html.Link().Href(path.CSSReset).Rel("stylesheet"),
			html.Style().TextUnsafe(
				pageStyle().String(),
			),
			flex.Styles(),
		),
		html.Body().Class(flex.VBox).Children(
			html.Div().Class(flex.HBox, headerBar, flex.CenterItems).Children(
				html.H1().Class(siteTitle).Text(title),
				html.Span().Class(flex.Grow),
			),
			html.Div().Class(flex.HBox, flex.Grow).Children(
				html.Div().Class(linksMenu, flex.VBox).Children(
					html.A().Class(menuItem, selectedMenuItem).Href(path.Videos.List).Text("Videos"),
					html.A().Class(menuItem).Href(path.Videos.Unencoded).Text("Processing"),
					html.A().Class(menuItem).Href(path.Files.List).Text("Files"),
					html.A().Class(menuItem).Href(path.Torrents).Text("Torrents"),
					html.A().Class(menuItem).Href(path.AddMagnetLink).Text("Add Magnet link"),
				),
				html.Div().Class(flex.Grow, mainSection, flex.VBox).Children(
					bodyContent...,
				),
			),
			html.Div().Class(flex.HBox, statusLine).Children(
				html.Span().Class(flex.Grow),
				html.Span().Class().Text(statusLineText),
			),
		),
	)
}
