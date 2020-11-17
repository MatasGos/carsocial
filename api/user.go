package api

import (
	"fmt"
	"net/http"
)

//PostUser create user object
func PostUser(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

//GetUser gets user object
func GetUser(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//PutUser updates user object
func PutUser(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//DeleteUser deletes user object
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("users")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//GetUserList returns users list
func GetUserList(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("users")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
