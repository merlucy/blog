package router

import (
	//"blog/model"
	"blog/model"
	"fmt"
	"html/template"
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

	r.ParseForm()

	if r.FormValue("title") == "" || r.FormValue("content") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var p, pc string
	p = "<div>"
	pc = "</div>"
	strs := []string{p, "", pc}

	s := strings.Split(r.FormValue("content"), "\n")

	for i, t := range s {
		if i%2 == 1 {
			continue
		}

		strs[1] = t
		s[i] = strings.Join(strs, "")
	}

	post := model.Post{
		Title: r.FormValue("title"),
		Body:  template.HTML(strings.Join(s, "")),
	}

	db.Create(&post)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	//fmt.Fprintf(w, "s with <p> is : %s", s)

}
