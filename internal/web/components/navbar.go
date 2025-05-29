package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Navbar(authState bool) Node {
	return Nav(Class("w-full max-w-screen-lg px-4 py-2 mx-auto lg:px-8 lg:py-3 mt-10"),
		Div(Class("container flex flex-wrap items-center justify-between mx-auto text-base font-semibold text-neutral-800"),
			A(Class("mr-4 block cursor-pointer py-1.5"), Href("/"), Text("DAVServer")),
			If(authState,
				A(Class("mr-4 block cursor-pointer py-1.5"), Href("/logout"), Text("Logout")),
			),
			If(authState == false,
				A(Class("mr-4 block cursor-pointer py-1.5"), Href("/login"), Text("Login")),
			),
		),
	)
}
