package main

import (
	"./shared"
	"github.com/go-humble/router"
	"github.com/steveoc64/formulate"
	"honnef.co/go/js/dom"
)

// ID    int    `db:"id"`
// Name  string `db:"name"`
// Descr string `db:"descr"`

func productList(context *router.Context) {
	go func() {

		data := []shared.Product{}
		err := restServer.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.ListForm{}
		form.New("fa-cube", "Product List")

		// Define the layout

		form.Column("Name", "Name")
		form.Column("Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/")
		})

		form.NewRowEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/product/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/product/" + key)
		})

		form.Render("prod-list", ".jass-main", &data)
	}()
}

func productEdit(context *router.Context) {
	go func() {

		data := shared.Product{}
		err := restServer.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.EditForm{}
		form.New("fa-cube", data.Name)

		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/products")
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

func productAdd(context *router.Context) {
	print("TODO - productAdd")
}
