package view

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/vidos/path"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/display"
	"github.com/sparkymat/webdsl/css/size"
)

const AppName = "Vidos"

const pageBody css.Class = "page-body"
const vbox css.Class = "vbox"
const hbox css.Class = "hbox"
const grow css.Class = "hbox"

func pageStyle() css.CssContainer {
	return css.CssContainer{}.Rules(
		css.AllSelectors(css.Body, css.Html).Style(
			css.Width(size.Percent(100)),
			css.Height(size.Percent(100)),
		),
		vbox.Style(
			css.Display(display.Flex),
			css.FlexDirection(css.Column),
		),
		hbox.Style(
			css.Display(display.Flex),
			css.FlexDirection(css.Row),
		),
		grow.Style(
			css.FlexGrow(1),
		),
	)
}

func Layout(bodyContent html.Node) html.Node {
	return html.Html().Children(
		html.Head().Children(
			//TODO do better here
			html.Title().Text(AppName),
			html.Link().Href(path.CssReset).Rel("stylesheet"),
			html.Style().Text(
				pageStyle().String(),
			),
		),
		html.Body().Children(
			html.Div().Class(vbox).Children(
				html.Div().Class(hbox).Children(
					html.H1().Text(AppName),
					html.Span().Class(grow),
					html.Span().Text("TODO, memory disk usage"),
				),
			),
			html.Div().Class(pageBody).Children(
				bodyContent,
			),
		),
	)
}
