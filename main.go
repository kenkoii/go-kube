package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World v2")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
