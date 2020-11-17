// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"net/http"
)

func postComment(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

func getComment(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func putComment(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func getCommentList(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("comment")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
