package router

import (
	"blog/middleware"
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
)

type NoteData struct {
	Button template.HTML
	Notes  []Note
}

type Note struct {
	Body           template.HTML
	ID             uint
	VisitorID      uint
	VisitorName    string
	VisitorProfile string
	CreatedAt      string
}

//NotesHandler fetches visiting notes in the database and renders Google log in button
//if the visitor is not signed in.
//
//Since the only user authorized to write blog posts is Admin, visitors require
//Google log in to write visiting notes.
//
//Need to handle for Facebook or github log in in the future.
func NotesHandler(w http.ResponseWriter, r *http.Request) {

	//Receive an instance of model.Visitor
	v := Visitor(r)

	fmt.Println(v)

	//Start a new db session
	db := Db(r)
	defer db.Commit()

	//Create an empty model.Note instance
	note := []model.Note{}

	//Find all notes from the database
	db.Find(&note)

	//Parse the template of the view
	t, err := Parse(noteList, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	//loginButton becomes a 'log in with Google' button if the visitor is not logged in
	//or it becomes a picture of the visitor
	var loginButton template.HTML
	var nd NoteData
	var ndd Note

	//v.ID is not 0 if the user is registered in the database, thus shows a button that
	//indicates that the visitor is logged in
	if v.ID != 0 {

		loginButton = template.HTML("<div class=\"container logincheck\"><img class=\"profile-pic\" src=\"" + v.Picture + "\"><div class=\"loginstatus\">You are temporarily signed in as " + v.Name + "</div></div>")

		//loginButton is copied to nd.Button, which in the end is passed through to
		//be rendered in a view
		nd.Button = loginButton

	} else {

		loginButton = "<a class=\"btn btn-primary\" href=\"/gologin\" role=\"button\">Sign-in with Google</a>"
		nd.Button = loginButton

	}

	//Order of the notes is reversed in order to present earlier notes on the top
	for i := len(note); i > 0; i-- {

		ndd = NoteConvert(&note[i-1])
		v2 := VisitorByID(int(ndd.VisitorID))
		ndd.VisitorName = v2.Name
		ndd.VisitorProfile = v2.Picture
		nd.Notes = append(nd.Notes, ndd)
	}

	t.Execute(w, nd)

}

func NoteConvert(note *model.Note) (n Note) {

	n = Note{
		Body:      note.Body,
		ID:        note.ID,
		VisitorID: note.VisitorID,
		CreatedAt: note.CreatedAt.Format("02 Jan 2006"),
	}
	return n
}

//Visitor() returns an instance of model.Visitor that matches the data of the
//visitor sent through the Cookie.
func Visitor(r *http.Request) model.Visitor {

	//Check if the visitor is logged in
	c, err := r.Cookie("VisitorEmail")

	//If the visitor is not logged in return an error message
	if err != nil {
		fmt.Println("No Visitor Email Cookie")
		return model.Visitor{}
	}

	e := c.Value

	db := middleware.Database

	v := model.Visitor{}
	db.Where("Email = ?", e).First(&v)

	if v.ID == 0 {
		fmt.Println("Visitor not registered")
		return model.Visitor{}
	}

	return v
}

func VisitorByID(id int) model.Visitor {

	db := middleware.Database

	v := model.Visitor{}
	db.Where("ID = ?", id).First(&v)

	if v.ID == 0 {
		fmt.Println("Visitor not registered")
		return model.Visitor{}
	}

	return v

}

func UploadNoteHandler(w http.ResponseWriter, r *http.Request) {

	v := Visitor(r)

	if v.ID == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db := Db(r)
	defer db.Commit()

	r.ParseForm()

	if r.FormValue("content") == "" {
		http.Redirect(w, r, "/visiting", http.StatusSeeOther)
		return
	}

	//paragraph takes a string as an argument and returns a summary and formatted content
	_, s := paragraph(r.FormValue("content"))

	note := model.Note{
		Body:      template.HTML(s),
		Visitor:   v,
		VisitorID: v.ID,
	}

	db.Create(&note)

	http.Redirect(w, r, "/visiting", http.StatusSeeOther)

}
