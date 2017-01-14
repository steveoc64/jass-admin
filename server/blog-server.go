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
	e.GET("/api/blog/:id", getBlog)
	e.PATCH("/api/blog/:id", updateBlog)
}

func getBlogs(c echo.Context) error {
	printLog(c)
	blogs := []shared.Blog{}

	err := DB.SQL(`select * from blog order by post_order desc`).QueryStructs(&blogs)
	if err != nil {
		println(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, blogs)
}

func getBlog(c echo.Context) error {
	printLog(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "")
	}
	blog := shared.Blog{}
	DB.SQL(`select * from blog where id=$1`, id).QueryStruct(&blog)
	return c.JSON(http.StatusOK, blog)
}

func updateBlog(c echo.Context) error {
	printLog(c)
	return c.JSON(http.StatusOK, "")
}
