package router

import (
	"github.com/gorilla/mux"
	"github.com/imjulianeral/mux-performance/controllers"
)

func SetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/json", controllers.ReturnJSON)

	return router
}
