package router

import (
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

//Constants to keep track of template file names
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

//BlogListHandler manages URL calls for blog list page
func BlogListHandler(w http.ResponseWriter, r *http.Request) {

	//DB instance
	db := Db(r)
	defer db.Commit()

	//Variable to hold the db search result of posts
	post := []model.Post{}

	//Find the last 5 posts recently created
	db.Order("created_at desc").Limit(5).Find(&post)

	//Parse the post list view page
	t, err := Parse(postList, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var pd PostData
	var p Post

	//Convert model.Post slice into router.Post slice using router.PostConvert
	for i := 0; i < len(post); i++ {
		p = PostConvert(&post[i])
		pd.Posts = append(pd.Posts, p)
	}

	//If the user is looged in, present the new post upload button
	if r.Context().Value("login") != nil {
		pd.Login = true
	}

	//Service the final blog list page
	t.Execute(w, pd)
}

//BlogPageHandler manages URL calls for a single blog page
func BlogPageHandler(w http.ResponseWriter, r *http.Request) {

	//Identify the post id by only stripping the relevant URL path
	id, _ := Id(r.URL.Path)

	db := Db(r)
	defer db.Commit()

	post := model.Post{}
	user := model.User{}
	db.First(&post, id)

	//Find the author(user) of the post
	db.Model(&post).Related(&user)

	fmt.Println("Found user:", user)

	//Find the post written by the indicated user
	db.Model(&user).Related(&post)

	fmt.Println("Found post written by user:", post)

	t, err := Parse(blogPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var p Post

	p = PostConvert(&post)
	fmt.Printf("		Show Post with ID: %d\n", p.ID)

	t.Execute(w, p)
}

//ProjectListHandler manages URL calls for a page listing all the projects
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

//ProjectPageHandler manages URL calls for a single project page
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

	var p Project

	p = ProjectConvert(&project)
	t.Execute(w, p)
}

//ProfileHandler manages URL calls for a the profile page
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my profile", r.URL.Path[1:])
}

//Parse function parses the multiple html files using template.ParseFiles
func Parse(url ...string) (t *template.Template, err error) {
	t, err = template.ParseFiles(url...)

	return t, err
}

//Id function identifies and returns the id part of the url
func Id(url string) (string, int) {
	params := strings.Split(url, "/")

	return params[len(params)-1], len(params)
}

//Returns the attached database session of the request.
func Db(r *http.Request) *gorm.DB {
	db := r.Context().Value("db")

	return db.(*gorm.DB)
}

//PostConvert converts the model.Post into router.Post
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

//ProjectConvert converts the model.Project into router.Project
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
