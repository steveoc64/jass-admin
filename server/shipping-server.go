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

	e.GET("/api/shipping", listShipping)
	e.POST("/api/shipping", addShipping)
	e.GET("/api/shipping/:id", getShipping)
	e.PATCH("/api/shipping/:id", updateShipping)
	e.DELETE("/api/shipping/:id", deleteShipping)

	e.GET("/api/region", listRegion)
	e.POST("/api/region", addRegion)
	e.GET("/api/region/:id", getRegion)
	e.PATCH("/api/region/:id", updateRegion)
	e.DELETE("/api/region/:id", deleteRegion)
}

func listShipping(c echo.Context) error {
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

func listRegion(c echo.Context) error {
	printLog(c)
	data := []shared.Region{}

	err := DB.SQL(`select * from region order by id`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getRegion(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.ShippingCode{}
	DB.SQL(`select * from region where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateRegion(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func addRegion(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteRegion(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
