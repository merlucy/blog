package middleware

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/jinzhu/gorm"

	"blog/model"
	"blog/router"
)

type Middleware func(http.Handler) http.Handler

func InitiateMiddleware(db *gorm.DB, mux http.Handler) http.Handler {

	//Middleware addition logic written here
	return Middlewares(db)(Logger(LoginCheck(mux)))

}

//Middlewares function returns a http.Handler that attaches the database session to the request.
//The returned http.Handler takes the next middleware as a parameter and calls the next middleware.
func Middlewares(db *gorm.DB) Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//Attach a database session to the request context.
			ctx := context.WithValue(r.Context(), "db", DBSession(db))

			//Cleans the path to remove unnecessary or useless parameters.
			//This is to prevent directing to a non-existent page
			rn := Clean(r)

			//Since WithContext returns a new request at a new address,
			//the new request is stored in the variable rn.
			rn = r.WithContext(ctx)

			//Calls the next middleware aka. handler
			next.ServeHTTP(w, rn)
		})
	}

}

//DBSession function returns a database session. Specifically, gorm.DB session.
func DBSession(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	return tx
}

//Simple logging middleware to keep track of requests made to the server.
func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("METHOD:%-5sURL:%s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}

//Clean function cleans the URL path of the request using path.Clean.
func Clean(r *http.Request) (rn *http.Request) {
	path := path.Clean(r.URL.Path)
	r.URL.Path = path
	rn = r
	return rn
}

func LoginCheck(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Print(r.Cookies())

		cs := r.Cookies()

		var e, p string
		var login bool
		var rn *http.Request

		if cs != nil {

			e = cs[0].Value
			p = cs[1].Value
			login = Login(e, p, r)
			fmt.Print(login)

			if login {

				ctx := context.WithValue(r.Context(), "login", "true")
				rn = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(w, rn)
	})

}

func Login(email, password string, r *http.Request) bool {
	db := router.Db(r)
	user := model.User{}
	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return false
	}

	if user.Email == email && user.Password == password {
		return true
	}

	return false
}
