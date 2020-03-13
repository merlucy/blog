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

	db := Db(r)
	defer db.Commit()

	tags := []model.Tag{}

	db.Find(&tags)

	var td TagData
	var tdd Tag

	for i := 0; i < len(tags); i++ {
		fmt.Println("Looping through for loop:", i)
		tdd = TagConvert(&tags[i])
		td.Tags = append(td.Tags, tdd)

	}

	t, err := Parse(uploadPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	t.Execute(w, td)
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

	fmt.Println("Form Value is ", r.FormValue("check2"))

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

//RemoveTags removes HTML tags in order to smooth the blog post body for user modification
//when the EditPost handler is called
func RemoveTags(c template.HTML) template.HTML {

	s := string(c)

	fmt.Println(s)

	s = strings.Replace(s, "<div class=\"pg\">", "\n", -1)
	s = strings.Replace(s, "</div>", "", -1)

	fmt.Println(s)

	return template.HTML(s)
}

func EditPageHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	db := Db(r)

	post := model.Post{}
	db.First(&post, id)

	t, err := Parse(editPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	post.Body = RemoveTags(post.Body)
	var p Post
	p = PostConvert(&post)

	t.Execute(w, p)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	db := model.DB
	post := model.Post{}

	r.ParseForm()

	if r.FormValue("title") == "" || r.FormValue("content") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db.Where("id = ?", id).Find(&post)
	summary, value := paragraph(r.FormValue("content"))

	post.Title = r.FormValue("title")
	post.Body = template.HTML(value)
	post.Summary = template.HTML(summary)

	db.Save(&post)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

//Need to handle for db sessions
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)

	model.DB.Delete(model.Post{}, "id = ?", id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
