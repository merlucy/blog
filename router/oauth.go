package router

import (
	"blog/middleware"
	"blog/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "779780041670-tlq08q3rio5kd551tnugildjfnj16cem.apps.googleusercontent.com",
		ClientSecret: "S8V_v0VqoGav6khlAdE0rhsZ",
		RedirectURL:  "http://www.yjinlee.com/gocb",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "jeffkim"
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {

	url := oauthConf.AuthCodeURL(oauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")

	if state != oauthStateString {

		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")

	token, err := oauthConf.Exchange(oauth2.NoContext, code)

	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" +
		url.QueryEscape(token.AccessToken))

	if err != nil {
		fmt.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//Creates visitor instance in the database.
	v := addVisitor(r, response)

	//Concatenates visitor log in info to the Cookie
	addVisitorCookie(w, r, v)

	log.Printf("parseResponseBody: %s\n", string(response))

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

//addVisitor unmarshals the response received from Google and creates new visitor
//instance in the database.
//
//Need to add logic for adding duplicate visitors
func addVisitor(r *http.Request, rsp []byte) model.Visitor {

	info := make(map[string]interface{})

	err := json.Unmarshal(rsp, &info)

	if err != nil {
		fmt.Print("Marshal error : ")
		fmt.Println(err)
		return model.Visitor{}
	}

	fmt.Println(info)

	v := model.Visitor{}

	if result, data := findDuplicateVisitor(info["email"].(string)); result == 1 {

		v = data
	} else {
		v = model.Visitor{
			Name:    info["name"].(string),
			Email:   info["email"].(string),
			Picture: info["picture"].(string),
		}

		db := Db(r)
		defer db.Commit()

		db.Create(&v)
		fmt.Printf("%v created", v)
	}

	return v
}

func findDuplicateVisitor(email string) (int, model.Visitor) {

	db := middleware.Database

	v := model.Visitor{}
	db.Where("Email = ?", email).First(&v)

	if v.ID != 0 {
		fmt.Println("Visitor is already registered in the database")
		return 1, model.Visitor{}
	}

	return 0, model.Visitor{}
}

//addVisitor Cookie appends visitor information to the Cookie for further interaction
//through Cookie during other requests.
//
//Need to add visitor logout function
func addVisitorCookie(w http.ResponseWriter, r *http.Request, v model.Visitor) {

	c := http.Cookie{
		Name:     "VisitorEmail",
		Value:    v.Email,
		HttpOnly: true,
	}

	http.SetCookie(w, &c)
}
