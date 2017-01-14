package main

import "honnef.co/go/js/dom"

func main() {
	w := dom.GetWindow()

	w.ScrollTo(0, 0)

	initResize()
	initRouter()
	initForms()
	// initBurger() - dont initburger till login

	// doLoginPage()
	showTopMenu()
}
