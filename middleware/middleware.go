package middleware

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
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
			rn := r.WithContext(ctx)
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

func PathParser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		p := []byte(r.URL.Path)
		l := len(r.URL.Path)

		if p[l-1] == '/' {
			next.ServeHTTP(w, r)
		} else {

			for i := l; i > 0; i-- {
				if p[i-1] == '/' {
					ctx := context.WithValue(r.Context(), "id", p[i:])

					//r.WithContext(ctx) returns a pointer to the new address of a copied request
					//Thus, we should either keep the pointer value if we want to play with it,
					//or we should use the result directly in the next function call to next.ServeHTTp
					r2 := r.WithContext(ctx)

					//Need to check if MUtex should be used

					p2 := make([]byte, 0, len(r2.URL.Path)+1)

					fmt.Println(p2)
					fmt.Printf("%x", ':')
					next.ServeHTTP(w, r2)
					break
				}
			}
		}
	})

}
