package middleware

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"blog/model"
	"github.com/jinzhu/gorm"
)

type Middleware func(http.Handler) http.Handler

var Database *gorm.DB

var middlewares = []Middleware{

	Logger,
	LoginCheck,
}

func InitiateMiddleware(db *gorm.DB, mux http.Handler) http.Handler {

	Database = db

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

		fmt.Println("Printing Cookies")
		fmt.Print(r.Cookies())

		c1, err1 := r.Cookie("Email")
		c2, err2 := r.Cookie("Password")
		var e, p string
		var login bool

		if err1 == nil && err2 == nil {

			var rn *http.Request
			e = c1.Value
			p = c2.Value
			fmt.Println("Authenticating")
			login = Login(e, p)
			fmt.Print(login)

			if login {
				fmt.Println("Login success")
				ctx := context.WithValue(r.Context(), "login", "true")
				fmt.Println(ctx)
				rn = r.WithContext(ctx)
				r = rn
			}

		}

		next.ServeHTTP(w, r)
	})

}

//Login function identifies the email and password with reference to the database.
//This does not use the database session created within the middleware to maintain
//the integrity of architecture.
//
//Further expansion of login server should be in mind, thus allowing for separation
//of login database call and handler's session database call.
func Login(email, password string) bool {

	fmt.Println("Login procedure")
	user := model.User{}
	fmt.Println(email)

	Database.Where("Email = ?", email).First(&user)
	fmt.Println("Okay")

	if user.ID == 0 {
		return false
	}

	if user.Email == email && user.Password == password {
		return true
	}

	return false
}
