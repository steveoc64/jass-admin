package main

import (
	"./shared"
	"github.com/go-humble/form"
	"honnef.co/go/js/dom"
)

func doLoginPage() {
	w := dom.GetWindow()
	doc := w.Document()

	loadTemplate("login", ".jass-login", &Session)
	fadeIn("jass-login")

	doc.QuerySelector("#l-loginbtn").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		print("click on login btn")

		username := doc.QuerySelector("#l-username").(*dom.HTMLInputElement).Value
		passwd := doc.QuerySelector("#l-passwd").(*dom.HTMLInputElement).Value
		print("username", username)
		print("passwd", passwd)

		loginCred := &shared.Login{}
		if err := form.Bind(loginCred); err != nil {
			print("bind error", err.Error())
		} else {
			print("login cred", loginCred)
		}

	})
}
