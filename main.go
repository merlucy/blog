package main

import (
	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
)

func main() {

	//Server is initiated by init() in server.go
	defer server.db.Close()

	log.Fatal(http.ListenAndServe(":8080", server.Server()))

}
