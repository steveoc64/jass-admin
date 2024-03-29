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

func initShipping() {
	Session.Router.HandleFunc("/shippings", shippingList)
	Session.Router.HandleFunc("/shipping/add", shippingAdd)
	Session.Router.HandleFunc("/shipping/{id}", shippingEdit)
}

func shippingList(context *router.Context) {
	go func() {

		data := []shared.ShippingCode{}
		err := restServer.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.ListForm{}
		form.New("fa-truck", "Shipping Types")

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
			Session.Navigate("/shipping/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/shipping/" + key)
		})

		form.Render("shipcode-list", ".jass-main", &data)
	}()
}

func shippingEdit(context *router.Context) {
	go func() {

		data := shared.ShippingCode{}
		err := restServer.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.EditForm{}
		form.New("fa-truck", data.Name)

		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/shippings")
		})

		form.DeleteEvent(func(evt dom.Event) {
			evt.PreventDefault()
			go func() {
				err := restServer.Delete(&data)
				if err != nil {
					print("REST delete", err.Error())
				} else {
					Session.Navigate("/shippings")
				}
			}()
		})

		form.SaveEvent(func(evt dom.Event) {
			evt.PreventDefault()
			go func() {
				form.Bind(&data)
				err := restServer.Update(&data)
				if err != nil {
					print("REST update", err.Error())
				} else {
					Session.Navigate("/shippings")
				}
			}()
		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}

func shippingAdd(context *router.Context) {
	go func() {

		data := shared.ShippingCode{}
		// Session.OrientationSensitive = true

		form := formulate.EditForm{}
		form.New("fa-truck", "New Shipping Code")

		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/shippings")
		})

		form.SaveEvent(func(evt dom.Event) {
			evt.PreventDefault()

			go func() {
				form.Bind(&data)
				err := restServer.Create(&data)
				if err != nil {
					print("REST create", err.Error())
				} else {
					// Session.Navigate(fmt.Sprintf("/shipping/%d", data.ID))
					Session.Navigate("/shippings")
				}
			}()
		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}
