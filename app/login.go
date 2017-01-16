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
	doc.QuerySelector("[name=username]").(*dom.HTMLInputElement).Focus()

	doc.QuerySelector("#sign-in").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()

		loginCred := &shared.Login{}
		f, err := form.Parse(doc.QuerySelector("form"))
		if err != nil {
			print("form parse error", err.Error())
		} else {
			f.Validate("username").Required()
			f.Validate("passwd").Required()
			if f.HasErrors() {
				for _, err := range f.Errors {
					print("err: ", err.Error())
				}
			} else {
				if err := f.Bind(loginCred); err != nil {
					print("bind error", err.Error())
				} else {
					go func() {
						print("login cred", loginCred)

						if err = apiServer.Create(loginCred); err != nil {
							print("REST err", err.Error())
						} else {
							print("UID = ", loginCred.UID)
							Session.SetUID(loginCred.UID)
							Session.Navigate("/")
						}
					}()
				}
			}
		}

	})
}
