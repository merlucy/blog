package main

import (
	"blog/middleware"
	"blog/router"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
	"net/http"
)

func main() {

	db, err := gorm.Open("mysql", "root:Gostanford1@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	mux := router.InitRouter()
	middle := middleware.Logger(mux)
	//parser := middleware.PathParser(middle)
	log.Fatal(http.ListenAndServe(":8080", middle))

}
