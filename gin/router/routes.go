package router

import (
	"github.com/gin-gonic/gin"
	"github.com/imjulianeral/gin-performance/controllers"
)

func SetRoutes(router *gin.Engine) {
	router.POST("/json", controllers.ReturnJSON)
}
