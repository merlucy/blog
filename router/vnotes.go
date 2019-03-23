package router

import (
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
)

type NoteData struct {
	Notes []Note
}

type Note struct {
	Body      template.HTML
	ID        uint
	CreatedAt string
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {

	db := Db(r)
	defer db.Commit()
	note := []model.Note{}
	db.Find(&note)

	t, err := Parse(noteList, header)
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

func NoteConvert(note *model.Note) (n Note) {

	n = Note{
		Body:      note.Body,
		ID:        note.ID,
		CreatedAt: note.CreatedAt.Format("02 Jan 2006"),
	}
	return n
}
