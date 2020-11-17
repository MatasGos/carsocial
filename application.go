// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//===================================

//===================================

//===================================

//---------------------------------------

func loadJson(title string) (string, error) {
	filename := "data/" + title + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func loadJsonList(title string) (string, error) {
	filename := "data/" + title + "s.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadJson(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", p)
}
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

//localhost:5000/view/FrontPage
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	flag.Parse()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("frontpage"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/cars", func(r chi.Router) {
		r.With(paginate).Get("/", getCarList) // GET /articles

		r.Post("/", postUser) // POST /articles

		// Subrouters:
		r.Route("/{carID}", func(r chi.Router) {
			r.Get("/", getCar)       // GET /articles/123
			r.Put("/", putCar)       // PUT /articles/123
			r.Delete("/", deleteCar) // DELETE /articles/123
		})
	})
	r.Route("/users", func(r chi.Router) {
		r.With(paginate).Get("/", getUserList) // GET /articles

		r.Post("/", putUser) // POST /articles

		// Subrouters:
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", getUser)       // GET /articles/123
			r.Put("/", putUser)       // PUT /articles/123
			r.Delete("/", deleteUser) // DELETE /articles/123
		})
	})
	r.Route("/posts", func(r chi.Router) {
		r.With(paginate).Get("/", getPostList) // GET /articles

		r.Post("/", postCar) // POST /articles

		// Subrouters:
		r.Route("/{postID}", func(r chi.Router) {
			r.Get("/", getPost)       // GET /articles/123
			r.Put("/", putPost)       // PUT /articles/123
			r.Delete("/", deletePost) // DELETE /articles/123
		})
	})
	r.Route("/comments", func(r chi.Router) {
		r.With(paginate).Get("/", getCommentList) // GET /articles

		r.Post("/", putComment) // POST /articles

		// Subrouters:
		r.Route("/{commentID}", func(r chi.Router) {
			r.Get("/", getComment)       // GET /articles/123
			r.Put("/", putComment)       // PUT /articles/123
			r.Delete("/", deleteComment) // DELETE /articles/123
		})
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
