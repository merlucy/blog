package middleware

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/jinzhu/gorm"
)

type Middleware func(http.Handler) http.Handler

func InitiateMiddleware(db *gorm.DB, mux http.Handler) http.Handler {

	//Middleware addition logic written here
	return Middlewares(db)(Logger(mux))

}

func Middlewares(db *gorm.DB) Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", DBSession(db))
			rn := Clean(r)
			rn = r.WithContext(ctx)
			next.ServeHTTP(w, rn)
		})
	}

}

func DBSession(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	return tx
}

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("METHOD:%-5sURL:%s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}

func Clean(r *http.Request) (rn *http.Request) {
	path := path.Clean(r.URL.Path)
	r.URL.Path = path
	rn = r
	return rn
}
