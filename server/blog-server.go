package main

import (
	// "fmt"
	"net/http"
	"strconv"

	"../shared"
	"github.com/labstack/echo"
)

func initBlogs(e *echo.Echo) {
	println("setting up routes for blogs")
	e.GET("/api/blog", getBlogs)
	e.POST("/api/blog/add", addBlog)
	e.GET("/api/blog/:id", getBlog)
	e.PATCH("/api/blog/:id", updateBlog)
	e.DELETE("/api/blog/:id", deleteBlog)
}

func getBlogs(c echo.Context) error {
	printLog(c)
	data := []shared.Blog{}

	err := DB.SQL(`select * from blog order by post_order desc`).QueryStructs(&data)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func getBlog(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	data := shared.Blog{}
	DB.SQL(`select * from blog where id=$1`, id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func updateBlog(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func addBlog(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}

func deleteBlog(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
