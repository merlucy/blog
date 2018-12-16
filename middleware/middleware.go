package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("METHOD:%-5sURL:%s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}

func PathParser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if string(r.URL.Path[len(r.URL.Path)-1]) == "/" {
			next.ServeHTTP(w, r)
		} else {

			for i := len(r.URL.Path); i > 0; i-- {
				if string(r.URL.Path[i-1]) == "/" {
					ctx := context.WithValue(r.Context(), "id", string(r.URL.Path[i:]))

					//r.WithContext(ctx) returns a pointer to the new address of a copied request
					//Thus, we should either keep the pointer value if we want to play with it,
					//or we should use the result directly in the next function call to next.ServeHTTp
					r2 := r.WithContext(ctx)
					//fmt.Println(r2.Context().Value("id"))
					next.ServeHTTP(w, r2)
					break
				}
			}
		}
	})

}
