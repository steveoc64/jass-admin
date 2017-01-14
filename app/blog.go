package main

import (
	"./shared"
	"github.com/go-humble/router"
	"github.com/steveoc64/formulate"
	"honnef.co/go/js/dom"
)

// ID              int       `db:"id"`
// PostOrder       int       `db:"post_order"`
// Image           string    `db:"image"`
// Date            time.Time `db:"date"`
// Name            string    `db:"name"`
// Title           string    `db:"title"`
// Content         string    `db:"content"`
// ShareTwitter    int       `db:"share_twitter"`
// ShareFacebook   int       `db:"share_facebook"`
// ShareInstagram  int       `db:"share_instagram"`
// ShareGooglePlus int       `db:"share_google_plus"`

func blogMaint(context *router.Context) {
	go func() {

		data := []shared.Blog{}
		err := conn.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		} else {
			print("got blogs", data)
		}

		form := formulate.ListForm{}
		form.New("fa-hashtag", "Blog List")

		// Define the layout

		form.Column("PostOrder", "Post Order")
		form.Column("Name", "Name")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/")
		})

		form.NewRowEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/blog/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/blog/" + key)
		})

		form.Render("blog-list", "jass-main", data)
	}()
}
