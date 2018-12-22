package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {

	temp, err := gorm.Open("mysql", "root:Gostanford1@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal()
	}

	DB = temp
	SetTables()
	fmt.Println("DONE SETTING TABLES")
	SetRelationship()
	fmt.Println("DONE SETTING RELATIONSHIPS")
	//Populate Tables
	//Populate(DB)
}

type User struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Email string
}

//Blog post
type Post struct {
	gorm.Model
	Title  string
	Body   string
	Summary string
	User   User //`gorm:"foreignkey:UserID"`
	UserID uint
}

type Project struct {
	gorm.Model
	Title  string
	Body   string
	Summary string
	User   User
	UserID uint
}

type Note struct {
	gorm.Model
	Body   string
	User   User
	UserID uint
}

/*
type Portfolio struct {


}
*/

func SetTables() {

	DB.CreateTable(&User{})
	DB.CreateTable(&Post{})
	DB.CreateTable(&Project{})
	DB.CreateTable(&Note{})

}

//Set Relationships
func SetRelationship() {

	var user User
	var post Post
	var project Project
	var note Note

	DB.Model(&user).Related(&post)
	DB.Model(&user).Related(&project)
	DB.Model(&user).Related(&note)

}
