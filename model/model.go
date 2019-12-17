package model

import (
	"fmt"
	"html/template"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {

	temp, err := gorm.Open("mysql", "root:Gostanford1!@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal()
	}

	DB = temp.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")

	SetTables()
	fmt.Println("DONE SETTING TABLES")
	//SetRelationship()
	fmt.Println("DONE SETTING RELATIONSHIPS")
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Post{})
	DB.AutoMigrate(&Project{})
	DB.AutoMigrate(&Visitor{})
	DB.AutoMigrate(&Note{})

	//Populate Tables
	Populate(DB)
}

//User struct which dictates what possessions users can have
type User struct {
	ID       uint `gorm:"primary_key"`
	Password string
	Name     string
	Email    string

	//User can have many posts
	Posts []Post `gorm:"foreignkey:UserID"`

	//User can have many projects
	Projects []Project `gorm:"foreignkey:UserID"`
}

//Visitor struct is created when a visitor logs in to the website using oauth
type Visitor struct {
	ID      uint `gorm:"primary_key`
	Name    string
	Email   string
	Picture string
	Link    string

	//Visitor can have many notes
	Notes []Note `gorm:"foreignkey:VisitorID"`
}

//Post struct is an object for blog post
type Post struct {
	gorm.Model
	Title   string
	Body    template.HTML `sql:"type:longtext"`
	Summary template.HTML `sql:"type:longtext:`

	//A blog post can have many tags
	Tags []Tag `gorm:"many2many:post_tags;"`

	//UserID as foreign key of the author
	UserID uint
}

//Tag struct is used to classify blog posts
type Tag struct {
	gorm.Model
	Name string

	//A tag can have many posts
	Posts []Post `gorm:"many2many:post_tags;"`
}

//Project struct contains information of user's project
type Project struct {
	gorm.Model
	Title   string
	Body    template.HTML `sql:"type:longtext"`
	Summary template.HTML

	//UserID as foreign key for the project owner
	UserID uint
}

//Note struct contains the contents of visiting notes by visitors
type Note struct {
	gorm.Model
	Body template.HTML

	//VisitorID as foreign key for the author
	VisitorID uint
}

//SetTables function creates all the tables for the database
func SetTables() {

	DB.CreateTable(&User{})
	DB.CreateTable(&Post{})
	DB.CreateTable(&Project{})
	DB.CreateTable(&Note{})
	DB.CreateTable(&Tag{})
}

/*
//Set Relationships
func SetRelationship() {

	var user User
	var post Post
	var project Project
	var note Note
	var visitor Visitor

		DB.Model(&user).Related(&post)
		DB.Model(&user).Related(&project)
		DB.Model(&visitor).Related(&note)
}*/
