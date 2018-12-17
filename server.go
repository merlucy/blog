package main

import (
	"blog/middleware"
	"blog/model"
	"blog/router"
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

	server.db = model.DB
}

func (s Server) Server() http.Handler {

	m := middleware.InitiateMiddleware(s.db, s.mux)
	return m
}
