package router

import (
	"blog/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {

}

func SigninPageHandler(w http.ResponseWriter, r *http.Request) {

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	
	db := Db(r)
	defer db.Commit()
	db.Create(&user)
	fmt.Printf("%v created", user)
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {

	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
		
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	
	db := Db(r)
	defer db.Commit()
	
	//Check if any element is empty
	
	//Implement find by username
	//db.First(&user, )
	
	//If no username, redirect
	//If there is a username, check password with the db.
	//If successful, redirect
	//Add session cookie, allowing for blog editing.
	
}
