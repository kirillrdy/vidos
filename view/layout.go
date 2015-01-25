package view

import "github.com/kirillrdy/nadeshiko/html"

func Layout(bodyContent html.Node) html.Node {
	return html.Html().Children(
		html.Head().Children(
			//TODO do better here
			html.Title().Text("Vidos"),
		),
		html.Body().Children(
			bodyContent,
		),
	)
}
