package main

import (
	// "fmt"
	"net/http"
	"strconv"

	"../shared"
	"github.com/labstack/echo"
)

func initProduct(e *echo.Echo) {
	println("setting up routes for products")
	e.GET("/api/category", listCategory)
	e.POST("/api/category", addCategory)
	e.GET("/api/category/:id", getCategory)
	e.PATCH("/api/category/:id", updateCategory)
	e.DELETE("/api/category/:id", deleteCategory)

	e.GET("/api/product", listProduct)
	e.POST("/api/product", addProduct)
	e.GET("/api/product/:id", getProduct)
	e.PATCH("/api/product/:id", updateProduct)
	e.DELETE("/api/product/:id", deleteProduct)
}

func listCategory(c echo.Context) error {
	printLog(c)
	data := []shared.Category{}

	err := DB.SQL(`select * from category order by name`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getCategory(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.Category{}
	DB.SQL(`select * from category where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateCategory(c echo.Context) error {
	data := shared.Category{}
	c.Bind(&data)
	printLog(c, data)

	_, err := DB.Update("category").
		SetWhitelist(data, "name", "descr").
		Where("id = $1", data.ID).
		Exec()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func addCategory(c echo.Context) error {
	data := shared.Category{}
	c.Bind(&data)
	printLog(c, data)

	id := 0
	err := DB.InsertInto("category").
		Whitelist("name", "descr").
		Record(data).
		Returning("id").
		QueryScalar(&id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	DB.SQL("select * from category where id=$1", id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func deleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	printLog(c, id)

	_, err := DB.SQL("delete from category where id=$1", id).Exec()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func listProduct(c echo.Context) error {
	printLog(c)
	data := []shared.Product{}

	err := DB.SQL(`select * from product order by name`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getProduct(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.Product{}
	DB.SQL(`select * from product where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func addProduct(c echo.Context) error {
	data := shared.Product{}
	c.Bind(&data)
	printLog(c, data)

	id := 0
	err := DB.InsertInto("product").
		Whitelist("name", "descr", "image",
			"volume_ml", "weight_g",
			"shipping_volume_ml", "shipping_weight_g",
			"shipping_code").
		Record(data).
		Returning("id").
		QueryScalar(&id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	DB.SQL("select * from product where id=$1", id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateProduct(c echo.Context) error {
	data := shared.Product{}
	c.Bind(&data)
	printLog(c, data)

	_, err := DB.Update("product").
		SetWhitelist(data, "name", "descr", "image",
			"volume_ml", "weight_g",
			"shipping_volume_ml", "shipping_weight_g",
			"shipping_code").
		Where("id = $1", data.ID).
		Exec()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	printLog(c, id)

	_, err := DB.SQL("delete from product where id=$1", id).Exec()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}
