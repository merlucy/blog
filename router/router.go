package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Route struct {
	Name    string
	Address string
}

var Routes = map[Route]http.HandlerFunc{
	Route{"IndexPage", "/"}:             IndexHandler,
	Route{"BlogList", "/blog"}:          BlogListHandler,
	Route{"BlogPage", "/blog/"}:         BlogPageHandler,
	Route{"ProjectList", "/projects"}:   ProjectListHandler,
	Route{"ProjectPage", "/projects/"}:  ProjectPageHandler,
	Route{"VisitingNotes", "/visiting"}: VisitingNotesHandler,
	Route{"Portfolio", "/Portfolio"}:    PortfolioHandler,
}

func InitRouter() *http.ServeMux {
	m := http.NewServeMux()

	for rt, h := range Routes {
		m.Handle(rt.Address, h)
	}
	return m
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HI", r.URL.Path[1:])
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
