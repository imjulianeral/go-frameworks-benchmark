package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imjulianeral/fiber-performance/controllers"
)

func SetRoutes(router *fiber.App) {
	router.Post("/json", controllers.ReturnJSON)
}
