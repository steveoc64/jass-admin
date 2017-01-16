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
	e.POST("/api/category/add", addCategory)
	e.GET("/api/category/:id", getCategory)
	e.PATCH("/api/category/:id", updateCategory)
	e.DELETE("/api/category/:id", deleteCategory)

	e.GET("/api/product", listProduct)
	e.POST("/api/product/new", addProduct)
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
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func addCategory(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteCategory(c echo.Context) error {
	printLog(c)
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
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func updateProduct(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteProduct(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
