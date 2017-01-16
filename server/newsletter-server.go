package main

import (
	// "fmt"
	"net/http"
	"strconv"

	"../shared"
	"github.com/labstack/echo"
)

func initNewsletters(e *echo.Echo) {
	println("setting up routes for newsletters")
	e.GET("/api/news", listNewsletter)
	e.POST("/api/news/add", addNewsletter)
	e.GET("/api/news/:id", getNewsletter)
	e.PATCH("/api/news/:id", updateNewsletter)
	e.DELETE("/api/news/:id", deleteNewsletter)
}

func listNewsletter(c echo.Context) error {
	printLog(c)
	data := []shared.Newsletter{}

	err := DB.SQL(`select * from newsletter order by date desc`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getNewsletter(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.SaleOrder{}
	DB.SQL(`select * from newsletter where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateNewsletter(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func addNewsletter(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteNewsletter(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
