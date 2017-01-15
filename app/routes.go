package main

import (
	"errors"
	"strconv"

	"github.com/go-humble/locstor"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func fixLinks() {
	Session.Router.InterceptLinks()
}

// Load a template and attach it to the specified element in the doc
func loadTemplate(template string, selector string, data interface{}) error {
	w := dom.GetWindow()
	doc := w.Document()

	t, err := GetTemplate(template)
	if t == nil {
		print("Failed to load template", template)
		return errors.New("Invalid template")
	}
	if err != nil {
		print(err.Error())
		return err
	}

	el := doc.QuerySelector(selector)
	// print("loadtemplate", template, "into", selector, "=", el)
	if el == nil {
		print("Could not find selector", selector)
		return errors.New("Invalid selector")
	}
	// print("looks ok adding template", t, "to", el, "with", data)
	if err := t.ExecuteEl(el, data); err != nil {
		print(err.Error())
		return err
	}
	Session.Router.InterceptLinks()
	return nil
}

func initRouter() {
	// print("initRouter")
	Session.Context = nil
	Session.ID = make(map[string]int)

	// Include public routes
	Session.Router = router.New()
	Session.Router.ShouldInterceptLinks = true
	Session.Router.HandleFunc("/", defaultRoute)
	Session.Router.HandleFunc("/blogs", blogList)
	Session.Router.HandleFunc("/blog/add", blogAdd)
	Session.Router.HandleFunc("/blog/{id}", blogEdit)
	Session.Router.HandleFunc("/categories", categoryList)
	Session.Router.HandleFunc("/category/add", categoryList)
	Session.Router.HandleFunc("/category/{id}", categoryEdit)
	Session.Router.HandleFunc("/products", productList)
	Session.Router.HandleFunc("/product/add", productAdd)
	Session.Router.HandleFunc("/product/{id}", productEdit)
	Session.Router.HandleFunc("/shippings", shippingList)
	Session.Router.HandleFunc("/shipping/add", shippingAdd)
	Session.Router.HandleFunc("/shipping/{id}", shippingEdit)
	Session.Router.HandleFunc("/orders", orderList)
	Session.Router.HandleFunc("/order/add", orderAdd)
	Session.Router.HandleFunc("/order/{id}", orderEdit)
	Session.Router.HandleFunc("/newsletters", newsList)
	Session.Router.HandleFunc("/news/add", newsAdd)
	Session.Router.HandleFunc("/news/{id}", newsEdit)
	Session.Router.Start()
}

func defaultRoute(context *router.Context) {
	// print("Nav to Default Route with UID", Session.UID)
	if Session.UID == 0 {
		uidstr, err := locstor.GetItem("uid")
		if err != nil {
			doLoginPage()
			return
		}

		uid, err := strconv.Atoi(uidstr)

		if uid != 0 {
			// print("loaded UID from local storage", uid)
			Session.UID = uid
			doMainPage()
		} else {
			// print("not logged in")
			doLoginPage()
		}
	} else {
		doMainPage()
	}
}
