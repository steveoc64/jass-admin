package main

import (
	// "fmt"
	"net/http"
	"strconv"

	"../shared"
	"github.com/labstack/echo"
)

func initShipping(e *echo.Echo) {
	println("setting up routes for shipping details")
	e.GET("/api/shipping", getShippings)
	e.POST("/api/shipping/add", addShipping)
	e.GET("/api/shipping/:id", getShipping)
	e.PATCH("/api/shipping/:id", updateShipping)
	e.DELETE("/api/shipping/:id", deleteShipping)
}

func getShippings(c echo.Context) error {
	printLog(c)
	data := []shared.ShippingCode{}

	err := DB.SQL(`select * from shipping_code order by id`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getShipping(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.ShippingCode{}
	DB.SQL(`select * from shipping_code where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateShipping(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func addShipping(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteShipping(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
