package router

import (
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

type PostData struct {
	Posts []Post
}

type Post struct {
	Title     string
	Body      string
	ID        uint
	CreatedAt string
}

type ProjectData struct {
	Projects []Project
}

type Project struct {
	Title     string
	Body      string
	ID        uint
	CreatedAt string
}

type NoteData struct {
	Notes []Note
}

type Note struct {
	Body      string
	ID        uint
	CreatedAt string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	post := []model.Post{}
	db.Find(&post)

	t, err := template.ParseFiles("templates/header.html")
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd PostData
	var pdd Post

	for _, p := range post {

		pdd = PostConvert(&p)
		pd.Posts = append(pd.Posts, pdd)
	}

	t.Execute(w, pd)
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	post := []model.Post{}
	db.Find(&post)

	t, err := template.ParseFiles("templates/header.html")
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd PostData
	var pdd Post

	for _, p := range post {

		pdd = PostConvert(&p)
		pd.Posts = append(pd.Posts, pdd)
	}

	t.Execute(w, pd)
}

func BlogPageHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	post := model.Post{}
	db := Db(r)
	db.First(&post, id)

	var pdd Post

	pdd = PostConvert(&post)

	fmt.Printf("ID Search Result: %d\n", pdd.ID)
}

func ProjectListHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	project := []model.Project{}
	db.Find(&project)

	t, err := template.ParseFiles("templates/projectList.html")
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd ProjectData
	var pdd Project

	for _, p := range project {

		pdd = ProjectConvert(&p)
		pd.Projects = append(pd.Projects, pdd)
	}

	t.Execute(w, pd)

}

func ProjectPageHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	project := model.Project{}
	db := Db(r)
	db.First(&project, id)

	var pdd Project

	pdd = ProjectConvert(&project)

	fmt.Printf("ID Search Result: %d\n", pdd.ID)

}

func VisitingNotesHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	note := []model.Note{}
	db.Find(&note)

	t, err := template.ParseFiles("templates/header.html")
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var nd NoteData
	var n Note

	for _, n := range note {

		n = NoteConvert(&note)
		nd.Notes = append(nd.Notes, n)
	}

	t.Execute(w, nd)

}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func Id(url string) (string, int) {

	params := strings.Split(url, "/")
	return params[len(params)-1], len(params)

}

func Db(r *http.Request) *gorm.DB {

	db := r.Context().Value("db")
	return db.(*gorm.DB)
}

func PostConvert(post *model.Post) (p Post) {

	p = Post{
		Title:     post.Title,
		Body:      post.Body,
		ID:        post.ID,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05 MST"),
	}

	return p
}

func ProjectConvert(project *model.Project) (p Project) {

	p = Project{
		Title:     project.Title,
		Body:      project.Body,
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format("2006-01-02 15:04:05 MST"),
	}

	return p
}

func ParamSame(subject, compare string) bool {
	return len(strings.Split(subject, "/")) == len(strings.Split(compare, "/"))
}

//Router needs
/*
* main page -> / -> show list of blog posts
* blogs -> blogs
* Intro -> introduction
* projects -> projects
* visitor page -> visiting
* Portfolio -> portfolio
 */
