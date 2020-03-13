package router

import (
	//	"blog/middleware"
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	header      = "templates/header.html"
	index       = "templates/index.html"
	postList    = "templates/bloglist.html"
	blogPage    = "templates/post.html"
	projectList = "templates/projectList.html"
	projectPage = "templates/project.html"
	noteList    = "templates/vnoteList.html"
	loginPage   = "templates/login.html"
	signupPage  = "templates/signup.html"
	signinPage  = "templates/signin.html"
	uploadPage  = "templates/upload.html"
	editPage    = "templates/edit.html"
)

type LoginInfo struct {
	Login bool
}

type PostData struct {
	Posts []Post
	Login bool
}

type Post struct {
	Title     string
	Body      template.HTML
	Summary   template.HTML
	ID        uint
	CreatedAt string
}

type ProjectData struct {
	Projects []Project
	Login    bool
}

type Project struct {
	Title     string
	Body      template.HTML
	Summary   template.HTML
	ID        uint
	CreatedAt string
}

func RemoveTags(c template.HTML) template.HTML {

	s := string(c)

	fmt.Println(s)

	s = strings.Replace(s, "<div class=\"pg\">", "\n", -1)
	s = strings.Replace(s, "</div>", "", -1)

	fmt.Println(s)

	return template.HTML(s)
}

//IndexHandler renders the index page which lists five latest blog posts.
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//Create a database session
	db := Db(r)
	defer db.Commit()

	//Create an empty model.Post slice
	post := []model.Post{}

	//Gorm SQL command to fetch five latest posts from the database
	db.Order("created_at desc").Limit(5).Find(&post)

	//Concatenate the index.html file with header.html file
	t, err := Parse(index, header)

	//If template parsing fails, print an error message and abort the request
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd PostData
	var pdd Post

	//Convert model.Post slice into router.Post slice using router.PostConvert
	for i := 0; i < len(post); i++ {
		pdd = PostConvert(&post[i])
		pd.Posts = append(pd.Posts, pdd)
	}

	//Service the final index page
	t.Execute(w, pd)
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	defer db.Commit()
	post := []model.Post{}
	db.Order("created_at desc").Limit(5).Find(&post)

	t, err := Parse(postList, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd PostData
	var pdd Post

	for i := 0; i < len(post); i++ {
		pdd = PostConvert(&post[i])
		pd.Posts = append(pd.Posts, pdd)
	}

	if r.Context().Value("login") != nil {
		pd.Login = true
	}

	t.Execute(w, pd)
}

func BlogPageHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	post := model.Post{}
	db := Db(r)
	defer db.Commit()
	db.First(&post, id)

	user := model.User{}

	//Find user who wrote this post
	db.Model(&post).Related(&user)

	fmt.Println("Found user:", user)

	db.Model(&user).Related(&post)

	fmt.Println("Found post written by user:", post)

	t, err := Parse(blogPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pdd Post

	pdd = PostConvert(&post)

	t.Execute(w, pdd)
	fmt.Printf("		Show Post with ID: %d\n", pdd.ID)
}

func ProjectListHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	defer db.Commit()
	project := []model.Project{}
	db.Find(&project)

	t, err := Parse(projectList, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd ProjectData
	var pdd Project

	for _, p := range project {

		pdd = ProjectConvert(&p)
		pd.Projects = append(pd.Projects, pdd)
	}

	if r.Context().Value("login") != nil {
		pd.Login = true
	}

	t.Execute(w, pd)
}

func ProjectPageHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := Id(r.URL.Path)
	project := model.Project{}
	db := Db(r)
	defer db.Commit()
	db.First(&project, id)

	t, err := Parse(projectPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pdd Project

	pdd = ProjectConvert(&project)
	t.Execute(w, pdd)
}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my profile", r.URL.Path[1:])
}

func Parse(url ...string) (t *template.Template, err error) {
	t, err = template.ParseFiles(url...)

	return t, err
}

func Id(url string) (string, int) {
	params := strings.Split(url, "/")

	return params[len(params)-1], len(params)
}

//Returns the attached database session of the request.
func Db(r *http.Request) *gorm.DB {
	db := r.Context().Value("db")

	return db.(*gorm.DB)
}

func PostConvert(post *model.Post) (p Post) {

	p = Post{
		Title:     post.Title,
		Body:      post.Body,
		Summary:   post.Summary,
		ID:        post.ID,
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

func ParamSame(subject, compare string) bool {
	return len(strings.Split(subject, "/")) == len(strings.Split(compare, "/"))
}
