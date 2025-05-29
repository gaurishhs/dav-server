package pages

import (
	"github.com/gaurishhs/dav-server/internal/web/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HomePage(authState bool) Node {
	return layouts.MainLayout("Home", authState,
		Section(Class("container mx-auto px-4 py-8"),
			H1(Class("text-3xl font-bold"), Text("Welcome to DAVServer!")),
			P(Class("mt-4"), Text("CalDAV/CardDAV server")),
		),
	)
}
