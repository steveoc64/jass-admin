package main

import "./shared"

func getItems() {
	Session.Items = []shared.Item{}
	GETJson("/api/items", &Session.Items)
}
