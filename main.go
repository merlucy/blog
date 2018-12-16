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
	parser := middleware.PathParser(middle)
	log.Fatal(http.ListenAndServe(":8080", parser))

}
