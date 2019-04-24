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

	summary, value := paragraph(r.FormValue("content"))

	post := model.Post{
		Title:   r.FormValue("title"),
		Body:    template.HTML(value),
		Summary: template.HTML(summary),
	}

	db.Create(&post)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

//summary function strips the first paragraph of the post.
func summary(c string) string {

	s := strings.Split(c, "\n")

	for i := 0; i < len(s); i++ {

		if s[i] == "" {
			continue
		} else {
			return s[i]
		}
	}

	fmt.Println("		Uploaded post has no content")

	return s[0]

}

func paragraph(c string) (sum, par string) {

	s := strings.Split(strings.Replace(c, "\r\n", "\n", -1), "\n")

	fmt.Printf("Content: %s", c)

	var p, pc string
	p = "<div class=\"pg\">"
	pc = "</div>"
	strs := []string{p, "", pc}

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

//Need to handle for db sessions
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)

	model.DB.Delete(model.Post{}, "id = ?", id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
