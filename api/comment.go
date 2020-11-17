package api

import (
	"fmt"
	"net/http"
)

//PostComment creates comment object
func PostComment(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

//GetComment returns comment object
func GetComment(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//PutComment updates comment object
func PutComment(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//DeleteComment removes comment object
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//GetCommentList returns comment list
func GetCommentList(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
