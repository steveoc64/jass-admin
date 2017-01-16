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

func initCategory() {
	Session.Router.HandleFunc("/categories", categoryList)
	Session.Router.HandleFunc("/category/add", categoryAdd)
	Session.Router.HandleFunc("/category/{id}", categoryEdit)

	Session.Router.HandleFunc("/products", productList)
	Session.Router.HandleFunc("/product/add", productAdd)
	Session.Router.HandleFunc("/product/{id}", productEdit)
}

func categoryList(context *router.Context) {
	go func() {

		data := []shared.Category{}
		err := restServer.ReadAll(&data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.ListForm{}
		form.New("fa-cubes", "Category List")

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
			Session.Navigate("/category/add")
		})

		form.RowEvent(func(key string) {
			Session.Navigate("/category/" + key)
		})

		form.Render("cat-list", ".jass-main", &data)
	}()
}

func categoryEdit(context *router.Context) {
	go func() {

		data := shared.Category{}
		err := restServer.Read(context.Params["id"], &data)
		if err != nil {
			print("REST error", err.Error())
			return
		}

		form := formulate.EditForm{}
		form.New("fa-cubes", data.Name)

		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/categories")
		})

		form.DeleteEvent(func(evt dom.Event) {
			evt.PreventDefault()
			go func() {
				err := restServer.Delete(&data)
				if err != nil {
					print("REST delete", err.Error())
				} else {
					Session.Navigate("/categories")
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
					Session.Navigate("/categories")
				}
			}()
		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}

func categoryAdd(context *router.Context) {
	print("add category")
	go func() {

		data := shared.Category{}
		form := formulate.EditForm{}
		form.New("fa-cubes", "Add New Category")

		form.Row(1).AddInput(1, "Name", "Name")
		form.Row(1).AddTextarea(1, "Description", "Descr")

		// Add event handlers
		form.CancelEvent(func(evt dom.Event) {
			evt.PreventDefault()
			Session.Navigate("/categories")
		})

		form.SaveEvent(func(evt dom.Event) {
			evt.PreventDefault()
			go func() {
				form.Bind(&data)
				err := restServer.Create(&data)
				if err != nil {
					print("REST create", err.Error())
				} else {
					// Session.Navigate(fmt.Sprintf("/category/%d", data.ID))
					Session.Navigate("/categories")
				}
			}()
		})

		form.Render("edit-form", ".jass-main", &data)
	}()
}
