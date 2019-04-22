package router

import (
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func UploadPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("login") == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

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
		return
	}

	value := r.FormValue("content")
	summary, paragraph := paragraph(value)

	post := model.Post{
		Title:   r.FormValue("title"),
		Body:    template.HTML(paragraph),
		Summary: template.HTML(summary),
	}

	db.Create(&post)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func paragraph(c string) (sum, par string) {

	fmt.Printf("Content: %s", c)

	var p, pc string
	p = "<div class=\"pg\">"
	pc = "</div>"
	strs := []string{p, "", pc}

	s := strings.Split(c, "\n")

	fmt.Printf("Split Text: %s", s)

	for i, t := range s {
		if i%2 == 1 {
			continue
		}

		strs[1] = t
		s[i] = strings.Join(strs, "")
	}

	fmt.Printf("Parsed String :%s", s)

	fmt.Printf("Summary: %s", s[0])
	return s[0], strings.Join(s, "")
}

/*
func EditPageHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	db := Db(r)

	post := model.Post{}

	db.First(&post, id)

	t, err := Parse("Edit.html", header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var p model.Post

	p = post

	t.Execute(w, p)

}
*/
//Need to handle for db sessions
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)

	model.DB.Delete(model.Post{}, "id = ?", id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
