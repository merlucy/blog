package router

import (
	"blog/model"
	//	"encoding/json"
	"fmt"
	"net/http"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {

	t, err := Parse(signupPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}
	t.Execute(w, nil)

}

func SigninPageHandler(w http.ResponseWriter, r *http.Request) {

	t, err := Parse(signinPage, header)
	if err != nil {
		fmt.Println("Template parse fail")
	}
	t.Execute(w, nil)

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Println("Decode complete")

	if r.FormValue("name") == "" || r.FormValue("email") == "" || r.FormValue("password") != r.FormValue("passwordcheck") {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}

	//need to add if there exists the same name in the database.

	user := model.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	fmt.Println(user)
	db := Db(r)
	defer db.Commit()
	db.Create(&user)
	fmt.Printf("%v created", user)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fmt.Println(r.Form)

	db := Db(r)
	defer db.Commit()

	var email, password string = r.FormValue("email"), r.FormValue("password")

	//Check if any element is empty
	if email == "" || password == "" {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}

	user := model.User{}

	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}

	fmt.Print(user)

	if user.Email == email && user.Password == password {

		c1 := http.Cookie{
			Name:     "Email",
			Value:    user.Email,
			HttpOnly: true,
		}

		c2 := http.Cookie{
			Name:     "Password",
			Value:    user.Password,
			HttpOnly: true,
		}

		http.SetCookie(w, &c1)
		http.SetCookie(w, &c2)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	//Implement find by username
	//db.First(&user, )

	//If no username, redirect
	//If there is a username, check password with the db.
	//If successful, redirect
	//Add session cookie, allowing for blog editing.

}

func SignoutHandler(w http.ResponseWriter, r *http.Request) {

	c1 := http.Cookie{
		Name:   "Email",
		MaxAge: -1,
	}

	c2 := http.Cookie{
		Name:   "Password",
		MaxAge: -1,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
