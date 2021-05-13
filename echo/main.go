package main

import (
	"github.com/imjulianeral/echo/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.SetRoutes(e)

	e.Logger.Fatal(e.Start(":4003"))
}
