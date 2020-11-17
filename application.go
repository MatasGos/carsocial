// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MatasGos/simple/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//=======================================

//---------------------------------------

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := api.LoadJSON(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", p)
}
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		r.With(paginate).Get("/", api.GetCarList) // GET /articles

		r.Post("/", api.PostUser) // POST /articles

		// Subrouters:
		r.Route("/{carID}", func(r chi.Router) {
			r.Get("/", api.GetCar)       // GET /articles/123
			r.Put("/", api.PutCar)       // PUT /articles/123
			r.Delete("/", api.DeleteCar) // DELETE /articles/123
		})
	})
	r.Route("/users", func(r chi.Router) {
		r.With(paginate).Get("/", api.GetUserList) // GET /articles

		r.Post("/", api.PutUser) // POST /articles

		// Subrouters:
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", api.GetUser)       // GET /articles/123
			r.Put("/", api.PutUser)       // PUT /articles/123
			r.Delete("/", api.DeleteUser) // DELETE /articles/123
		})
	})
	r.Route("/posts", func(r chi.Router) {
		r.With(paginate).Get("/", api.GetPostList) // GET /articles

		r.Post("/", api.PostCar) // POST /articles

		// Subrouters:
		r.Route("/{postID}", func(r chi.Router) {
			r.Get("/", api.GetPost)       // GET /articles/123
			r.Put("/", api.PutPost)       // PUT /articles/123
			r.Delete("/", api.DeletePost) // DELETE /articles/123
		})
	})
	r.Route("/comments", func(r chi.Router) {
		r.With(paginate).Get("/", api.GetCommentList) // GET /articles

		r.Post("/", api.PutComment) // POST /articles

		// Subrouters:
		r.Route("/{commentID}", func(r chi.Router) {
			r.Get("/", api.GetComment)       // GET /articles/123
			r.Put("/", api.PutComment)       // PUT /articles/123
			r.Delete("/", api.DeleteComment) // DELETE /articles/123
		})
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
