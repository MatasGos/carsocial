package api

import (
	"fmt"
	"net/http"
)

//PostPost create post object
func PostPost(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("post")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

//GetPost gets post object
func GetPost(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("post")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//PutPost updates post object
func PutPost(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("post")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//DeletePost delets post object
func DeletePost(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("posts")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//GetPostList gets posts list
func GetPostList(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("posts")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
