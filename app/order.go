package main

import (
	"fmt"

	"./shared"
	"github.com/go-humble/router"
	"github.com/steveoc64/formulate"
	"honnef.co/go/js/dom"
)

// ID    int    `db:"id"`
// Name  string `db:"name"`
// Descr string `db:"descr"`

func orderList(context *router.Context) {
	go func() {

		data := []shared.SaleOrder{}
		err := restServer.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.ListForm{}
		form.New("fa-vcard", "Order List")

		// Define the layout

		form.DateColumn("Date", "Date")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/")
		})

		form.NewRowEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/order/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/order/" + key)
		})

		form.Render("order-list", ".jass-main", &data)
	}()
}

func orderEdit(context *router.Context) {
	go func() {

		data := shared.SaleOrder{}
		err := restServer.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.EditForm{}
		form.New("fa-vcard", fmt.Sprintf("Order #%06d", data.ID))

		form.Row(1).AddDate(1, "Date", "Date")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/orders")
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

func orderAdd(context *router.Context) {
	print("TODO - productAdd")
}
