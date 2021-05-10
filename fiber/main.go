package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imjulianeral/fiber-performance/router"
)

func main() {
	app := fiber.New()

	router.SetRoutes(app)

	app.Listen(":3000")
}
