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

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"

	"github.com/MatasGos/simple/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
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
	api.OpenDatabase()
	defer api.CloseDatabase()

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

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(api.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(isAdmin)
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(api.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/comments", func(r chi.Router) {
			r.With(paginate).Get("/", api.GetCommentList)

			r.Post("/", api.PostComment)

			// Subrouters:
			r.Route("/{commentID}", func(r chi.Router) {
				r.Get("/", api.GetComment)
				r.Put("/", api.PutComment)
				r.Delete("/", api.DeleteComment)
			})
		})
		r.Route("/cars", func(r chi.Router) {
			r.With(paginate).Get("/", api.GetCarList)

			r.Post("/", api.PostCar)
			// Subrouters:
			r.Route("/{carID}", func(r chi.Router) {
				r.Get("/", api.GetCar)
				r.Put("/", api.PutCar)
				r.Delete("/", api.DeleteCar)
			})
		})
		r.Route("/users", func(r chi.Router) {
			r.With(paginate).Get("/", api.GetUserList)
			r.Post("/", api.PostUser)

			// Subrouters:
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", api.GetUser)
				r.Put("/", api.PutUser)
				r.Delete("/", api.DeleteUser)
			})
		})
		r.Route("/posts", func(r chi.Router) {
			r.With(paginate).Get("/", api.GetPostList)

			r.Post("/", api.PostPost)
			// Subrouters:
			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", api.GetPost)
				r.Put("/", api.PutPost)
				r.Delete("/", api.DeletePost)
			})
		})
	})
	r.Group(func(r chi.Router) {
		r.Get("/comments/{commentID}", api.GetComment)
		r.Get("/comments/", api.GetCommentList)

		r.Get("/cars/{carID}", api.GetCar)
		r.Get("/cars/", api.GetCarList)

		r.Get("/posts/{postID}", api.GetPost)
		r.Get("/posts/", api.GetPost)

		r.Post("/users/", api.PostUser)

	})
	r.Route("/login", func(r chi.Router) {
		r.Post("/", api.Login)
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func init() {
	api.TokenAuth = jwtauth.New("HS256", []byte(api.Jwtsecret), nil)
}

func isAdmin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := api.TokenAuth.Decode(jwtauth.TokenFromHeader(r))
		if err != nil {
			http.Error(w, "Not authorize", http.StatusUnauthorized)
		}
		fmt.Println(token.Claims.(jwt.MapClaims)["role"])
		if token.Claims.(jwt.MapClaims)["role"] != "admin" {
			http.Error(w, "Not authorize", http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
