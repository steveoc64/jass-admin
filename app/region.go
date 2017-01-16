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

func regionList(context *router.Context) {
	go func() {

		data := []shared.Region{}
		err := restServer.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.ListForm{}
		form.New("fa-globe", "Shipping Regions")

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
			Session.Navigate("/region/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/region/" + key)
		})

		form.Render("region-list", ".jass-main", &data)
	}()
}

func regionEdit(context *router.Context) {
	go func() {

		data := shared.Region{}
		err := restServer.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.EditForm{}
		form.New("fa-globe", data.Name)

		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/regions")
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

func regionAdd(context *router.Context) {
	print("TODO - regionAdd")
}
