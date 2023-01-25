package main

import (
	"os"
	"strings"

	"github.com/hoondal6970/learngo/learngo/scrapper"
	"github.com/labstack/echo"
)

const file_NAME string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(file_NAME)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(file_NAME, file_NAME)
}

func main() {
	scrapper.Scrape("term")
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1234"))
}
