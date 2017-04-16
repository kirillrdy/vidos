package flex

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/display"
	"github.com/sparkymat/webdsl/css/flex"
	"github.com/sparkymat/webdsl/css/size"
)

const (
	VBox        css.Class = "vbox"
	HBox        css.Class = "hbox"
	Wrap        css.Class = "wrap"
	Grow        css.Class = "grow"
	NoGrow      css.Class = "no-grow"
	CenterItems css.Class = "align-items-center"
)

func Styles() html.Node {
	style := css.Stylesheet(
		css.AllSelectors(css.Body, css.Html).Style(
			css.Width(size.Percent(100)),
			css.Height(size.Percent(100)),
		),
		VBox.Style(
			css.Display(display.Flex),
			css.FlexDirection(flex.Column),
		),
		HBox.Style(
			css.Display(display.Flex),
			css.FlexDirection(flex.Row),
		),
		Wrap.Style(
			css.FlexWrap(flex.Wrap),
			//TODO technically this doesn't belong here, but i think it will only ever be used here
			css.AlignContent(css.FlexStart),
		),
		Grow.Style(
			css.FlexGrow(1),
		),
		NoGrow.Style(
			css.FlexShrink(0),
		),
		CenterItems.Style(
			css.AlignItems(css.Center),
		),
	)
	return html.Style().Text(
		style.String(),
	)
}
