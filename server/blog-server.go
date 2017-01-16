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
	e.GET("/api/blog", listBlog)
	e.POST("/api/blog", addBlog)
	e.GET("/api/blog/:id", getBlog)
	e.PATCH("/api/blog/:id", updateBlog)
	e.DELETE("/api/blog/:id", deleteBlog)
}

func listBlog(c echo.Context) error {
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
	data := shared.Blog{}
	c.Bind(&data)
	printLog(c, data)

	_, err := DB.Update("blog").
		SetWhitelist(data, "post_order", "date", "name", "title", "content", "archived").
		Where("id = $1", data.ID).
		Exec()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func addBlog(c echo.Context) error {
	data := shared.Blog{}
	c.Bind(&data)
	printLog(c, data)

	// Get a new post order
	DB.SQL("select max(post_order)+10 from blog").QueryScalar(&data.PostOrder)

	id := 0
	err := DB.InsertInto("blog").
		Whitelist("post_order", "date", "name", "title", "content", "archived").
		Record(data).
		Returning("id").
		QueryScalar(&id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}
	DB.SQL("select * from blog where id=$1", id).QueryStruct(&data)
	return c.JSON(http.StatusOK, data)
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	printLog(c, id)

	_, err := DB.SQL("delete from blog where id=$1", id).Exec()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
