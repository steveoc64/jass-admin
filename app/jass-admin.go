package main

import (
	"github.com/go-humble/rest"
	"github.com/steveoc64/formulate"
	"honnef.co/go/js/dom"
)

var codeVersion = "0.1.0"
var restServer = &rest.Client{ContentType: rest.ContentJSON}

func main() {
	w := dom.GetWindow()
	w.ScrollTo(0, 0)

	formulate.Templates(GetTemplate)

	print("Booting up jass-admin -", codeVersion)

	initResize()
	initRouter()
	initForms()
	// initBurger() - dont initburger till login

	// doLoginPage()
	showTopMenu()
}
