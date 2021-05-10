package main

import (
	"log"
	"net/http"

	"github.com/imjulianeral/mux-performance/router"
)

func main() {
	srv := &http.Server{
		Handler: router.SetRoutes(),
		Addr:    "localhost:4001",
	}

	log.Fatal(srv.ListenAndServe())
}
