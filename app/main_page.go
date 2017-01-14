package main

func doMainPage() {
	print("loading main-page")
	loadTemplate("main-page", ".jass-main", &Session)
	print("done")
	fadeIn("jass-main")
}
