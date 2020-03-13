package router

import (
	"blog/middleware"
	"blog/model"
	"fmt"
	"net/http"
)

const (
	tagPage = "templates/tags.html"
)

type TagData struct {
	Tags []Tag
}

type Tag struct {
	ID   uint
	Name string
}

func TagHandler(w http.ResponseWriter, r *http.Request) {

	//	printf("Entering
	//Create a database session
	db := middleware.Database

	//Concatenate the index.html file with header.html file
	t, err := Parse(tagPage, header)

	//If template parsing fails, print an error message and abort the request
	if err != nil {
		fmt.Println("Template parse fail")
	}

	//Create an empty model.Post slice
	fmt.Println("Creating Tag")
	tag := []model.Tag{}
	post := []model.Post{}
	db.First(&post)
	fmt.Println("First Tag is ", tag)

	//Searches for tags based on given post
	db.Model(&post).Related(&tag, "Tags")

	fmt.Println(tag)

	var td TagData
	var tdd Tag

	//Convert model.Post slice into router.Post slice using router.PostConvert
	for i := 0; i < len(tag); i++ {
		tdd = TagConvert(&tag[i])
		td.Tags = append(td.Tags, tdd)
	}

	//Service the final index page
	t.Execute(w, td)
}

func TagConvert(tag *model.Tag) (t Tag) {
	t = Tag{
		Name: tag.Name,
		ID:   tag.ID,
	}

	return t
}
