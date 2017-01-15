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

func blogList(context *router.Context) {
	go func() {

		data := []shared.Blog{}
		err := conn.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
			// } else {
			// 	print("got blogs", data)
		}

		form := formulate.ListForm{}
		form.New("fa-hashtag", "Blog List")

		// Define the layout

		form.BoolColumn("Archived", "Archived")
		form.Column("Seq", "PostOrder")
		form.ImgColumn("Image", "GetImageURL")
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

		form.Render("blog-list", ".jass-main", &data)
	}()
}

func blogEdit(context *router.Context) {
	go func() {

		data := shared.Blog{}
		err := conn.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		} else {
			print("got blog", data)
			data.ImageURL = data.GetImageURL()
		}

		Session.MobileSensitive = true
		Session.OrientationSensitive = true

		form := formulate.EditForm{}
		form.New("fa-hashtag", data.Name)

		if Session.Mobile() {
			form.Row(1).AddNumber(1, "Seq", "PostOrder", "0")
			form.Row(1).AddCheck(1, "Archived", "Archived")
		} else {
			form.Row(2).AddNumber(1, "Seq", "PostOrder", "0").AddCheck(1, "Archived", "Archived")
		}
		form.Row(1).AddImage(1, "Image", "ImageURL")
		form.Row(1).AddInput(1, "Short Title", "Name")
		form.Row(1).AddInput(1, "Title", "Title")
		form.Row(1).AddTextarea(1, "Content", "Content")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/blogs")
		})

		form.SaveEvent(func(evt dom.Event) {
			evt.PreventDefault()
			form.Bind(&data)
			print("post bind", data)

			// rest.Update(data)

		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}

func blogAdd(context *router.Context) {
	print("TODO - add blog")
}
