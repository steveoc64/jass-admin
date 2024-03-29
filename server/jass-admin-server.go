package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/steveoc64/godev/db"
	runner "gopkg.in/mgutz/dat.v1/sqlx-runner"
)

var e *echo.Echo
var DB *runner.DB

func printLog(c echo.Context, s ...interface{}) {
	r := c.Request()
	realIP := r.Header["X-Forwarded-For"]
	theIP := r.RemoteAddr
	if len(realIP) > 0 {
		theIP = realIP[0]
	}
	log.Println(theIP, r.Method, r.URL.Path, s)
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	Config = LoadConfig()

	// Start up the basic web server
	e = echo.New()
	e.Debug = true

	e.Static("/", "public")

	initLogin(e)
	initBlogs(e)
	initProduct(e)
	initShipping(e)
	initOrders(e)
	initNewsletters(e)

	// Start up the mail server
	if Config.MailServer == "" {
		println("Mail is turned OFF")
	} else {
		println("Mail server =", Config.MailServer)

		MailChannel = InitMailer(Config)
	}

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			// errorCode := httpError.Code()
			errorCode := httpError.Code
			switch errorCode {
			case http.StatusNotFound:
				// TODO handle not found case
				// log.Println("Not Found", err.Error())
				// We are usually here due to an F5 refresh, in which case
				// the URL is not expected to be there
				context.Redirect(http.StatusMovedPermanently, "/")
			default:
				// TODO handle any other case
			}
		}
	}

	e.Use(middleware.Recover())
	// e.Use(middleware.Gzip())
	e.Debug = Config.Debug
	// echocors.Init(e, Config.Debug)

	// Connect to the database
	DB = db.Init(Config.DataSourceName)

	// Start the web server
	if Config.Debug {
		log.Printf("... Starting Jass Admin Server on port %d", Config.WebPort)
	}

	errRun := e.Start(fmt.Sprintf(":%d", Config.WebPort))
	if errRun != nil {
		println("Error: ", errRun.Error())
	}
	println("Jass Admin Server All Done")

}
