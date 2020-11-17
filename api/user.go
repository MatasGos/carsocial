package api

import (
	"fmt"
	"net/http"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("users")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("users")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
