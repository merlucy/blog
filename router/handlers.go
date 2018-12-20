package router

import (
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	index       = "templates/header.html"
	postList    = "templates/header.html"
	projectList = "templates/projectList.html"
	noteList    = "templates/vnoteList.html"
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
	defer db.Commit()
	post := []model.Post{}
	db.Find(&post)

	t, err := template.ParseFiles(index)
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
	defer db.Commit()
	post := []model.Post{}
	db.Find(&post)

	t, err := template.ParseFiles(postList)
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
	defer db.Commit()
	db.First(&post, id)

	var pdd Post

	pdd = PostConvert(&post)

	fmt.Printf("ID Search Result: %d\n", pdd.ID)
}

func ProjectListHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	defer db.Commit()
	project := []model.Project{}
	db.Find(&project)

	t, err := template.ParseFiles(projectList)
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
	defer db.Commit()
	db.First(&project, id)

	var pdd Project

	pdd = ProjectConvert(&project)

	fmt.Printf("ID Search Result: %d\n", pdd.ID)

}

func NotesHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	defer db.Commit()
	note := []model.Note{}
	db.Find(&note)

	t, err := template.ParseFiles(noteList)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var nd NoteData
	var ndd Note

	for _, n := range note {

		ndd = NoteConvert(&n)
		nd.Notes = append(nd.Notes, ndd)
	}

	t.Execute(w, nd)

}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my profile", r.URL.Path[1:])
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
		Title: post.Title,
		Body:  post.Body,
		ID:    post.ID,

		CreatedAt: post.CreatedAt.Format("02 Jan 2006"),
	}

	return p
}

func ProjectConvert(project *model.Project) (p Project) {

	p = Project{
		Title:     project.Title,
		Body:      project.Body,
		ID:        project.ID,
		CreatedAt: project.CreatedAt.Format("02 Jan 2006"),
	}

	return p
}

func NoteConvert(note *model.Note) (n Note) {

	n = Note{
		Body:      note.Body,
		ID:        note.ID,
		CreatedAt: note.CreatedAt.Format("02 Jan 2006"),
	}
	return n
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
