package main

import (
	"net/http"
	"strings"

	"../shared"

	"github.com/labstack/echo"
)

func initLogin(e *echo.Echo) {
	e.POST("/api/login", doLogin)
}

func doLogin(c echo.Context) error {

	loginCred := &shared.Login{}
	if err := c.Bind(loginCred); err != nil {
		println("bind error", err.Error())
	} else {
		printLog(c, *loginCred)
		// fmt.Printf("got login cred %v\n", *loginCred)
		switch strings.ToLower(loginCred.Username) {
		case "steve":
			if loginCred.Passwd == "unx911zxx" {
				loginCred.UID = 1
			}
		case "kal":
			if loginCred.Passwd == "!Gordon5045" {
				loginCred.UID = 2

			}
		case "kat":
			if loginCred.Passwd == "fysherdog775" {
				loginCred.UID = 3
			}
		default:
			loginCred.Username = ""
			loginCred.Passwd = ""
			loginCred.Result = "Invalid username or password"
		}
		return c.JSON(http.StatusOK, loginCred)
	}
	return nil
}
