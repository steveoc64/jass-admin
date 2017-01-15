package main

import (
	"./shared"
	"github.com/go-humble/router"
	"github.com/steveoc64/formulate"
	"honnef.co/go/js/dom"
)

// ID      int       `db:"id"`
// Date    time.Time `db:"date"`
// Name    string    `db:"name"`
// Content string    `db:"content"`

func newsList(context *router.Context) {
	go func() {

		data := []shared.Newsletter{}
		err := conn.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.ListForm{}
		form.New("fa-newspaper", "Newsletters")

		// Define the layout

		form.DateColumn("Date", "Date")
		form.Column("Name", "Name")
		form.Column("Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/")
		})

		form.NewRowEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/news/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/news/" + key)
		})

		form.Render("prod-list", ".jass-main", &data)
	}()
}

func newsEdit(context *router.Context) {
	go func() {

		data := shared.Newsletter{}
		err := conn.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.EditForm{}
		form.New("fa-newspaper", data.Name)

		form.Row(1).AddDate(1, "Date", "Date")
		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Content", "Content")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/newsletters")
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

func newsAdd(context *router.Context) {
	print("TODO - newsAdd")
}
