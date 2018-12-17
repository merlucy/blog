package main

import (
	"blog/middleware"
	"blog/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Server struct {
	mux *http.ServeMux
	db  *gorm.DB
}

var server Server

func init() {
	server.mux = router.InitRouter()

	db, err := gorm.Open("mysql", "root:Gostanford1@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal()
	}

	server.db = db
}

func (s Server) Server() http.Handler {

	m := middleware.InitiateMiddleware(s.db, s.mux)
	return m
}
