package web

import (
	"fmt"
	"runtime"
	"syscall"

	"github.com/kirillrdy/vidos/flex"
	"github.com/kirillrdy/vidos/util"
	"github.com/kirillrdy/web"
	reset "github.com/kirillrdy/web/css"
	"github.com/kirillrdy/web/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/font"
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
			css.FontWeight(font.WeightBold),
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

	//TODO use gohuman for formatting numbers
	// note that forceful type casting necessary here due to different types on different platforms (linux vs FreeBSD)
	freeStorage := float64(uint64(fsStat.Bavail)*uint64(fsStat.Bsize)) / float64(1024*1024*1024)
	memoryUsed := float64(memStat.Alloc) / float64(1024*1024)

	return fmt.Sprintf("OS:%v/%v FreeStorage:%.2f Gb, MemUsed: %.2f Mb", runtime.GOOS, runtime.GOARCH, freeStorage, memoryUsed)
}

//ToHTML returns html content for a given page
//TODO fix regression with overflow in the main grow section
func (page Page) ToHTML(bodyContent ...html.Node) html.Node {

	pageTitle := fmt.Sprintf("%v - %v", page.application.Name, page.Title)

	var links []html.Node
	for _, item := range page.application.Menu {
		links = append(links, html.A().Class(menuItem, selectedMenuItem).Href(item.Path).Text(item.Title))
	}

	return html.Html().Children(
		html.Head().Children(
			html.Title().Text(pageTitle),
			html.Link().Href(reset.ResetCSSPath).Rel("stylesheet"),
			html.Style().TextUnsafe(
				pageStyle().String(),
			),
			flex.Styles(),
		),
		html.Body().Class(flex.VBox).Children(
			html.Div().Class(flex.HBox, headerBar, flex.CenterItems).Children(
				html.H1().Class(siteTitle).Text(page.application.Name),
				html.Span().Class(flex.Grow),
			),
			html.Div().Class(flex.HBox, flex.Grow).Children(
				html.Div().Class(linksMenu, flex.VBox).Children(
					links...,
				),
				html.Div().Class(flex.Grow, mainSection, flex.VBox).Children(
					bodyContent...,
				),
			),
			html.Div().Class(flex.HBox, statusLine).Children(
				html.Span().Class(flex.Grow),
				html.Span().Class().Text(statusLineText()),
			),
		),
	)
}

// Page represents a single web page
type Page struct {
	application Application
	Path        web.Path
	Title       string
}

// Application represents a web application
// TODO application could have a list of default middleware
type Application struct {
	Name string
	Menu []Page
}

// NewPage creates a new page for a given application
// TODO possibly also register handler or something like that ?
func (application Application) NewPage(title string, path web.Path) Page {
	return Page{application: application, Path: path, Title: title}
}
