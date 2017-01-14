package main

func doMainPage() {
	loadTemplate("main-page", ".jass-main", &Session)
	fadeIn("jass-main")
	handleActionGridClick()
}
