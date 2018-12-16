package main

import (
	"blog/middleware"
	"blog/router"
	"log"
	"net/http"
)

func main() {

	mux := router.InitRouter()
	middle := middleware.Logger(mux)
	log.Fatal(http.ListenAndServe(":8080", middle))

}
