package main

import (
	"blog/middleware"
	"blog/model"
	"blog/router"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//type Server contains *http.ServeMux and *gorm.DB database instance address.
type Server struct {
	mux *http.ServeMux
	db  *gorm.DB
}

var server Server

//Initiates the server by calling model.DB and router.InitRouter().
func init() {

	//server.mux is initiated by router.InitRouter().
	//The method to initialize *http.ServeMux instance is in the router package.
	server.mux = router.InitRouter()

	//Database instance is copied from the inialized db instance from 'model' package.
	server.db = model.DB
}

//Initiates middleware by calling middleware.InitiateMiddleware(*gorm.DB, *http.ServeMux).
func (s Server) Server() http.Handler {

	m := middleware.InitiateMiddleware(s.db, s.mux)
	return m
}
