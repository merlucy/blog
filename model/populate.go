package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Populate(db *gorm.DB) {

	Users(db)
	Posts(db)
	Projects(db)
	Notes(db)
	Tags(db)
}

func Users(db *gorm.DB) {
	fmt.Println("SETTING USERS")
	examples := []User{
		User{Name: "Andy", Email: "Andy@test.com", Password: "1234"},
		User{Name: "Brian", Email: "Brian@test.com", Password: "1234"},
		User{Name: "Joe", Email: "Joe@test.com", Password: "1234"},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

func Posts(db *gorm.DB) {

	fmt.Println("SETTING POSTS")
	examples := []Post{
		Post{Title: "LOL", Body: "LOLBA", Summary: "Hi", Tag: Tag{TagID: 1, Name: "First Tag"}, UserID: 1},
		Post{Title: "LUL", Body: "LULBA", UserID: 2},
		Post{Title: "LIL", Body: "LILBA", UserID: 3},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

func Projects(db *gorm.DB) {

	fmt.Println("SETTING PROJECTS")
	examples := []Project{
		Project{Title: "Jeff", Body: "LOLBA", UserID: 1},
		Project{Title: "Daniel", Body: "LULBA", UserID: 2},
		Project{Title: "Wanggui", Body: "LILBA", UserID: 3},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

func Tags(db *gorm.DB) {
	fmt.Println("SETTING TAGS")
	examples := []Tag{
		Tag{Name: "First Tag"},
		Tag{Name: "Second Tag"},
		Tag{Name: "Third Tag"},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

func Notes(db *gorm.DB) {
	fmt.Println("SETTING NOTES")
	examples := []Note{
		Note{Body: "LOLBA", VisitorID: 1},
		Note{Body: "LULBA", VisitorID: 2},
		Note{Body: "LILBA", VisitorID: 3},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}
