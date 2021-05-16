package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imjulianeral/gin-performance/router"
)

func main() {
	app := gin.Default()

	router.SetRoutes(app)

	app.Run(":4004")
}
