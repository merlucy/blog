package router

import (
	"blog/middleware"
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type NoteData struct {
	Button template.HTML
	Notes  []Note
}

type Note struct {
	Body      template.HTML
	ID        uint
	CreatedAt string
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {

	v := Visitor(r)

	fmt.Println(v)

	db := Db(r)
	defer db.Commit()
	note := []model.Note{}
	db.Find(&note)

	t, err := Parse(noteList, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var loginButton template.HTML
	var nd NoteData
	var ndd Note

	if v.ID != 0 {

		loginButton = "<div class=\"container loginstatus\"><img src=\"" + v.Link + "><div>You are temporarily signed in as " + v.Name + "</div></div>"

		nd.Button = loginButton
	} else {
		loginButton = "<a class=\"btn btn-primary\" href=\"/gologin\" role=\"button\">Sign-in with Google</a>"
		nd.Button = loginButton
	}

	for _, n := range note {

		ndd = NoteConvert(&n)
		nd.Notes = append(nd.Notes, ndd)
	}

	t.Execute(w, nd)

}

func NoteConvert(note *model.Note) (n Note) {

	n = Note{
		Body:      note.Body,
		ID:        note.ID,
		CreatedAt: note.CreatedAt.Format("02 Jan 2006"),
	}
	return n
}

func Visitor(r *http.Request) model.Visitor {

	c, err := r.Cookie("VisitorEmail")

	if err != nil {

		fmt.Println("No Cookie")
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

func UploadNoteHandler(w http.ResponseWriter, r *http.Request) {

	v := Visitor(r)

	if v.ID == 0 {
		return
	}

	db := Db(r)
	defer db.Commit()

	r.ParseForm()

	if r.FormValue("content") == "" {
		http.Redirect(w, r, "/visiting", http.StatusSeeOther)
	}

	var p, pc string
	p = "<div class=\"pg\">"
	fmt.Println(p)
	pc = "</div>"
	strs := []string{p, "", pc}

	s := strings.Split(r.FormValue("content"), "\n")

	for i, t := range s {
		if i%2 == 1 {
			continue
		}

		strs[1] = t
		s[i] = strings.Join(strs, "")
	}

	note := model.Note{
		Body:      template.HTML(strings.Join(s, "")),
		Visitor:   v,
		VisitorID: v.ID,
	}

	db.Create(&note)

	http.Redirect(w, r, "/visiting", http.StatusSeeOther)

}
