package main

import (
	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
)

//Server starts by initializing the server and database instances.
func main() {

	defer server.db.Close()

	//Server is initiated by init() in server.go by calling server.Server().
	log.Fatal(http.ListenAndServe(":8080", server.Server()))

}
