package main

import (
	"net/http"

	"github.com/hoondal6970/learngo/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	scrapper.Scrape("term")
	e := echo.New()
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":1234"))
}
