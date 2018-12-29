package router

import (
	//"blog/model"
	"fmt"
	"net/http"
	"strings"
)

func UploadPageHandler(w http.ResponseWriter, r *http.Request) {

	t, err := Parse(uploadPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}
	t.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	defer db.Commit()
	//post := []model.Post{}

	r.ParseForm()

	s := strings.Split(r.FormValue("content"), "\n")
	fmt.Println(s)

	if r.FormValue("title") == "" || r.FormValue("content") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
