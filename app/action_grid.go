package main

import (
	"honnef.co/go/js/dom"
)

func handleActionGridClick() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector(".action-grid").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		url := getElementURL(evt.Target())
		if url != "" {
			Session.Navigate(url)
		}
	})
}

func getElementURL(el dom.Element) string {
	if el == nil {
		return ""
	}
	retval := el.GetAttribute("url")
	if retval != "" {
		return retval
	}

	// we have no URL, so step up to our parent
	parentElement := el.ParentElement()
	return getElementURL(parentElement)
}
