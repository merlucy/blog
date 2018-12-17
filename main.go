package main

import (
	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
)

func main() {

	defer server.db.Close()

	//	middle := middleware.Logger(server.mux)
	log.Fatal(http.ListenAndServe(":8080", server.Server()))

}
