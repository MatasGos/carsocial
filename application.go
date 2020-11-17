// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

<<<<<<< HEAD
=======
//var validPath = regexp.MustCompile("^/(user|comment|car|post)/([a-zA-Z0-9]+)?$")
var validPath = regexp.MustCompile("^/(user|comment|car|post)/([0-9]+)?$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Print(m)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		if m[2] == "" {
			m[1] = m[1] + "s"
		}
		fn(w, r, m[1])
	}
}
func carHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	if len(m) > 2 {
		if m[2] == "" {
			if r.Method == "GET" {
				getCarList(w, r)
				return
			} else if r.Method == "POST" {
				postCar(w, r)
				return
			}
		} else {
			if r.Method == "GET" {
				getCar(w, r)
				return
			} else if r.Method == "PUT" {
				putCar(w, r)
				return
			} else if r.Method == "DELETE" {
				deleteCar(w, r)
				return
			}
		}
	}
	http.NotFound(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	if len(m) > 2 {
		if m[2] == "" {
			if r.Method == "GET" {
				getPostList(w, r)
				return
			} else if r.Method == "POST" {
				postPost(w, r)
				return
			}
		} else {
			if r.Method == "GET" {
				getPost(w, r)
				return
			} else if r.Method == "PUT" {
				putPost(w, r)
				return
			} else if r.Method == "DELETE" {
				deletePost(w, r)
				return
			}
		}
	}
	http.NotFound(w, r)
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	if len(m) > 2 {
		if m[2] == "" {
			if r.Method == "GET" {
				getCommentList(w, r)
				return
			} else if r.Method == "POST" {
				postComment(w, r)
				return
			}
		} else {
			if r.Method == "GET" {
				getComment(w, r)
				return
			} else if r.Method == "PUT" {
				putComment(w, r)
				return
			} else if r.Method == "DELETE" {
				deleteComment(w, r)
				return
			}
		}
	}
	http.NotFound(w, r)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	if len(m) > 2 {
		if m[2] == "" {
			if r.Method == "GET" {
				getUserList(w, r)
				return
			} else if r.Method == "POST" {
				postUser(w, r)
				return
			}
		} else {
			if r.Method == "GET" {
				getUser(w, r)
				return
			} else if r.Method == "PUT" {
				putUser(w, r)
				return
			} else if r.Method == "DELETE" {
				deleteUser(w, r)
				return
			}
		}
	}

	http.NotFound(w, r)
}
func postCar(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("car")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

func getCar(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("car")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func putCar(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("car")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("cars")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

func getCarList(w http.ResponseWriter, r *http.Request) {
	p, err := loadJson("cars")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

>>>>>>> parent of 35cef41... added go-chi
//===================================

//===================================

//===================================

//---------------------------------------

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadJson(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", p)
}

//localhost:8080/view/FrontPage
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
<<<<<<< HEAD
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
=======
	http.HandleFunc("/car/", carHandler)
	http.HandleFunc("/comment/", commentHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/post/", postHandler)

	//http.HandleFunc("/", frontHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
>>>>>>> parent of 35cef41... added go-chi
}
