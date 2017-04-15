package view

import (
	"fmt"
	"log"
	"runtime"
	"syscall"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/layout"
	"github.com/kirillrdy/vidos/path"
	"github.com/kirillrdy/vidos/util"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/overflow"
	"github.com/sparkymat/webdsl/css/size"
	"github.com/sparkymat/webdsl/css/textdecoration"
)

const appName = "Видос"

var padding = size.Px(10)

const siteTitle css.Class = "site-title"
const linksMenu css.Class = "links-menu"
const centerItems css.Class = "align-items-center"
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
		centerItems.Style(
			css.AlignItems(css.Center),
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
	if err != nil {
		log.Print(err)
	}

	freeStorage := float64(uint64(fsStat.Bavail)*fsStat.Bsize) / float64(1024*1024*1024)
	memoryUsed := float64(memStat.Alloc) / float64(1024*1024)

	return fmt.Sprintf("OS:%v/%v FreeStorage:%.2f Gb, MemUsed: %.2f Mb", runtime.GOOS, runtime.GOARCH, freeStorage, memoryUsed)
}

//Page returns the main layout of the application
//TODO fix regression with overflow in the main grow section
func Page(title string, bodyContent ...html.Node) html.Node {

	title = fmt.Sprintf("%v - %v", appName, title)

	statusLineText := statusLineText()
	return html.Html().Children(
		html.Head().Children(
			html.Title().Text(title),
			html.Link().Href(path.CSSReset).Rel("stylesheet"),
			html.Style().TextUnsafe(
				pageStyle().String(),
			),
			layout.Styles(),
		),
		html.Body().Class(layout.VBox).Children(
			html.Div().Class(layout.HBox, headerBar, centerItems).Children(
				html.H1().Class(siteTitle).Text(appName),
				html.Span().Class(layout.Grow),
			),
			html.Div().Class(layout.HBox, layout.Grow).Children(
				html.Div().Class(linksMenu, layout.VBox).Children(
					html.A().Class(menuItem, selectedMenuItem).Href(path.Videos.List).Text("Videos"),
					html.A().Class(menuItem).Href(path.Videos.Unencoded).Text("Processing"),
					html.A().Class(menuItem).Href(path.Files.List).Text("Files"),
					html.A().Class(menuItem).Href(path.Torrents).Text("Torrents"),
					html.A().Class(menuItem).Href(path.AddMagnetLink).Text("Add Magnet link"),
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
