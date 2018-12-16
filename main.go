package main

import (
	"blog/router"
	"log"
	"net/http"
)

func main() {

	mux := router.InitRouter()

	log.Fatal(http.ListenAndServe(":8080", mux))

}
