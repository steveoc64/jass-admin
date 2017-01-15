package main

import (
	// "fmt"
	"net/http"
	"strconv"

	"../shared"
	"github.com/labstack/echo"
)

func initOrders(e *echo.Echo) {
	println("setting up routes for order details")
	e.GET("/api/order", getOrders)
	e.POST("/api/order/add", addOrder)
	e.GET("/api/order/:id", getOrder)
	e.PATCH("/api/order/:id", updateOrder)
	e.DELETE("/api/order/:id", deleteOrder)
}

func getOrders(c echo.Context) error {
	printLog(c)
	data := []shared.SaleOrder{}

	err := DB.SQL(`select * from sale_order order by date desc`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getOrder(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.SaleOrder{}
	DB.SQL(`select * from sale_order where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateOrder(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func addOrder(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteOrder(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
