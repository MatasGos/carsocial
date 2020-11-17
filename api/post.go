package api

import (
	"fmt"
	"net/http"
)

func postPost(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("post")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("post")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func putPost(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("post")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("posts")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func getPostList(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("posts")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
