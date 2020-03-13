package router

import (
	"blog/middleware"
	"blog/model"
	"fmt"
	"net/http"
)

const (
	catEditPage = "templates/catedit.html"
)

type Category struct {
	ID   uint
	Name string
}

type CategoryData struct {
	Categories []Category
	Login      bool
}

func CategoryConvert(category *model.Category) (c Category) {
	c = Category{
		Name: category.Name,
		ID:   category.ID,
	}
	return c
}

func CategoryEditPageHandler(w http.ResponseWriter, r *http.Request) {

	cs := []model.Category{}
	db := middleware.Database

	db.Find(&cs)
	fmt.Println(cs)

	t, err := Parse(catEditPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}

	var cd CategoryData
	var c Category

	for i := 0; i < len(cs); i++ {
		c = CategoryConvert(&cs[i])
		cd.Categories = append(cd.Categories, c)
	}

	t.Execute(w, cd)
}

func CategoryEditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Category delete called")

	db := model.DB
	c := model.Category{}
	r.ParseForm()

	if r.FormValue("name") == "" {
		fmt.Println("Redirecting")
		http.Redirect(w, r, "/catedit", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	fmt.Println(name)

	db.Where("name = ?", name).Find(&c)
	fmt.Println(c)

	db.Delete(&c)
	http.Redirect(w, r, "/catedit", http.StatusSeeOther)
}
