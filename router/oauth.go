package router

import (
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

	addVisitor(response)

	log.Printf("parseResponseBody: %s\n", string(response))

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func addVisitor(rsp []byte) {

	info := make(map[string]interface{})

	err := json.Unmarshal(rsp, &info)

	if err != nil {
		fmt.Println("Marshal error")
		fmt.Println(err)
		return
	}

	fmt.Println(info)
}
