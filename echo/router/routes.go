package router

import (
	"github.com/imjulianeral/echo/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutes(router *echo.Echo) {
	router.POST("/json", controllers.ReturnJSON)

}
