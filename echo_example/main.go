package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	l := log.New("echo")
	l.DisableColor()
	l.SetLevel(log.INFO)

	e := echo.New()
	e.Logger = l
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"_": "echo server example",
		})
	})
	e.GET("/plain", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!\n")
	})
	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello World!</h1>\n")
	})
	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	e.Logger.Info("Listen on 0.0.0.0:8000, see http://127.0.0.1:8000")
	e.Logger.Fatal(e.Start(":8000"))
}
