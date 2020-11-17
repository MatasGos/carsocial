package api

import (
	"fmt"
	"net/http"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("user")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("users")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("users")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
