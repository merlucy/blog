package router

import (
	"net/http"
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
	Route{"VisitingNotes", "/visiting"}: NotesHandler,
	Route{"Portfolio", "/portfolio"}:    PortfolioHandler,
	Route{"Profile", "/profile"}:        ProfileHandler,
	Route{"SigninPage", "/signin"}:      SigninPageHandler,
	Route{"SignupPage", "/signup"}:      SignupPageHandler,
	Route{"Signin", "/signing"}:         SigninHandler,
	Route{"Signup", "/sup"}:             SignupHandler,
	Route{"Signout", "/signout"}:        SignoutHandler,
	Route{"Upload", "/upload"}:          UploadPageHandler,
	Route{"UploadPage", "/upl"}:         UploadHandler,
	Route{"GoogleLogin", "/gologin"}:    GoogleLoginHandler,
	Route{"GoogleCallback", "/gocb"}:    GoogleCallbackHandler,
	Route{"UploadNote", "/uplnote"}:     UploadNoteHandler,
}

func InitRouter() *http.ServeMux {
	m := http.NewServeMux()

	for rt, h := range Routes {
		m.Handle(rt.Address, h)
	}

	fs := http.FileServer(http.Dir("templates"))
	m.Handle("/templates/", http.StripPrefix("/templates/", fs))
	return m
}
