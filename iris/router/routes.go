package router

import (
	"github.com/imjulianeral/iris-performance/controllers"
	"github.com/kataras/iris/v12"
)

func SetRoutes(app *iris.Application) {
	app.Post("/json", controllers.ReturnJSON)
}
