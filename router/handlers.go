package router

import (
	"blog/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	db := r.Context().Value("db")
	//fmt.Println(db)
	post := model.Post{}
	result := db.(*gorm.DB).Find(&post)

	fmt.Printf("Result is :%vi\n", post)
	fmt.Printf("HEX: %v\n", result)
	t, err := template.ParseFiles("header.html")
	if err != nil {
		fmt.Println("Template parse fail")
	}
	var Title string
	Title = "Brian"
	t.Execute(w, Title)
	//fmt.Fprintf(w, "HI", r.URL.Path[1:])
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello", r.URL.Path[1:])
}

func BlogPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ID is :", Id(r.URL.Path))
	fmt.Fprintf(w, "POST!!!!", r.URL.Path[1:])
}

func ProjectListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func ProjectPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func VisitingNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page", r.URL.Path[1:])
}

func Id(url string) string {

	params := strings.Split(url, "/")
	return params[len(params)-1]

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
