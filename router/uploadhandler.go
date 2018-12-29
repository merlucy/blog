package router

import (
	"fmt"
	"net/http"
)

func UploadPageHandler(w http.ResponseWriter, r *http.Request) {

	t, err := Parse(uploadPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}
	t.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
