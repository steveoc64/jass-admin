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
		form.ImgColumn("Image", "GetImageURL")
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
		cats := []shared.Category{}
		restServer.ReadAll(&cats)
		data.ImageURL = data.GetImageURL()

		form := formulate.EditForm{}
		form.New("fa-cube", data.Name)

		form.Row(1).AddSelect(1, "Category", "CatID", cats, "ID", "Name", 1, 0)
		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddInput(1, "Image", "Image")
		form.Row(1).AddImage(1, "Image Preview", "ImageURL")

		Session.MobileSensitive = true
		if Session.SubMobile() {
			form.Row(1).AddNumber(1, "Volume ml", "VolumeML", "1")
			form.Row(1).AddNumber(1, "Weight g", "WeightG", "1")
			form.Row(1).AddNumber(1, "Shipping Volume ml", "ShippingVolumeML", "1")
			form.Row(1).AddNumber(1, "Shipping Weight g", "ShippingWeightG", "1")
		} else {
			form.Row(2).
				AddNumber(1, "Volume ml", "VolumeML", "1").
				AddNumber(1, "Weight g", "WeightG", "1")
			form.Row(2).
				AddNumber(1, "Shipping Volume ml", "ShippingVolumeML", "1").
				AddNumber(1, "Shipping Weight g", "ShippingWeightG", "1")
		}
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/products")
		})

		form.DeleteEvent(func(evt dom.Event) {
			evt.PreventDefault()
			go func() {
				err := restServer.Delete(&data)
				if err != nil {
					print("REST delete", err.Error())
				} else {
					Session.Navigate("/products")
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
					Session.Navigate("/products")
				}
			}()
		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}

func productAdd(context *router.Context) {
	go func() {

		data := shared.Product{}
		cats := []shared.Category{}
		restServer.ReadAll(&cats)

		form := formulate.EditForm{}
		form.New("fa-cube", "Add New Product")

		form.Row(1).AddSelect(1, "Category", "CatID", cats, "ID", "Name", 1, 0)
		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/products")
		})

		form.SaveEvent(func(evt dom.Event) {
			evt.PreventDefault()
			go func() {
				form.Bind(&data)
				err := restServer.Create(&data)
				if err != nil {
					print("REST create", err.Error())
				} else {
					Session.Navigate(fmt.Sprintf("/product/%d", data.ID))
				}
			}()
		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}
