package main

import (
	"net/http"

	"github.com/imjulianeral/iris-performance/router"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	router.SetRoutes(app)

	srv := &http.Server{Addr: ":4002"}
	app.Run(iris.Server(srv))
}
