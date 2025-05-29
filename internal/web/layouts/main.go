package layouts

import (
	"github.com/gaurishhs/dav-server/internal/web/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func MainLayout(title string, authState bool, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(
				Rel("stylesheet"),
				Href("/assets/main.css"),
			),
		},
		Body: []Node{
			Class("text-neutral-800 bg-neutral-100"),
			components.Navbar(authState),
			Group(children),
		},
	})
}
