// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

import (
    "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)
//var validPath = regexp.MustCompile("^/(user|comment|car|post)/([a-zA-Z0-9]+)?$")
var validPath = regexp.MustCompile("^/(user|comment|car|post)/([0-9]+)?$")

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        fmt.Print(m)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        if(m[2]==""){
            m[1]=m[1]+"s"
        }
        fn(w, r, m[1])
	}
}
func carHandler(w http.ResponseWriter, r *http.Request){
    m := validPath.FindStringSubmatch(r.URL.Path)
    w.Header().Set("Content-Type", "application/json")

    if(len(m)>2){
    if(m[2]==""){
        if(r.Method =="GET"){
            getCarList(w, r)
            return
        } else if(r.Method == "POST"){
            postCar(w,r)
            return
        }  
    } else {
        if(r.Method =="GET"){
            getCar(w, r)
            return
        } else if(r.Method == "PUT"){
            putCar(w,r)
            return
        } else if(r.Method == "DELETE"){
            deleteCar(w,r)
            return
        }  
    }
}
http.NotFound(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request){
    m := validPath.FindStringSubmatch(r.URL.Path)
    w.Header().Set("Content-Type", "application/json")

    if(len(m)>2){
    if(m[2]==""){
        if(r.Method =="GET"){
            getPostList(w, r)
            return
        } else if(r.Method == "POST"){
            postPost(w,r)
            return
        }  
    } else {
        if(r.Method =="GET"){
            getPost(w, r)
            return
        } else if(r.Method == "PUT"){
            putPost(w,r)
            return
        } else if(r.Method == "DELETE"){
            deletePost(w,r)
            return
        }  
    }
}
http.NotFound(w, r)
}

func commentHandler(w http.ResponseWriter, r *http.Request){
    m := validPath.FindStringSubmatch(r.URL.Path)
    w.Header().Set("Content-Type", "application/json")

    if(len(m)>2){
    if(m[2]==""){
        if(r.Method =="GET"){
            getCommentList(w, r)
            return
        } else if(r.Method == "POST"){
            postComment(w,r)
            return
        }  
    } else {
        if(r.Method =="GET"){
            getComment(w, r)
            return
        } else if(r.Method == "PUT"){
            putComment(w,r)
            return
        } else if(r.Method == "DELETE"){
            deleteComment(w,r)
            return
        }  
    }
}
http.NotFound(w, r)
}


func userHandler(w http.ResponseWriter, r *http.Request){
    m := validPath.FindStringSubmatch(r.URL.Path)
    w.Header().Set("Content-Type", "application/json")

    if(len(m)>2){
    if(m[2]==""){
        if(r.Method =="GET"){
            getUserList(w, r)
            return
        } else if(r.Method == "POST"){
            postUser(w,r)
            return
        }  
    } else {
        if(r.Method =="GET"){
            getUser(w, r)
            return
        } else if(r.Method == "PUT"){
            putUser(w,r)
            return
        } else if(r.Method == "DELETE"){
            deleteUser(w,r)
            return
        }  
    }
}

    http.NotFound(w, r)
}
func postCar(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("car")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", p)
}

func getCar(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("car")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func putCar(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("car")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func deleteCar(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("cars")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func getCarList(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("cars")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

//===================================
func postPost(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("post")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", p)
}

func getPost(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("post")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func putPost(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("post")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func deletePost(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("posts")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func getPostList(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("posts")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

//===================================
func postUser(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("user")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", p)
}

func getUser(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("user")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func putUser(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("user")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("users")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func getUserList(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("users")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}
//===================================
func postComment(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("comment")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", p)
}

func getComment(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("comment")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func putComment(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("comment")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func deleteComment(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("comment")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

func getCommentList(w http.ResponseWriter, r *http.Request){
    p, err := loadJson("comment")
	if err != nil {
        http.NotFound(w, r)
        fmt.Print(err)
        return
    }  
    fmt.Fprintf(w, "%s", p)
}

//---------------------------------------

func loadJson(title string) (string, error) {
	filename := "data/"+title + ".json"
	body, err := ioutil.ReadFile(filename)
    if err != nil {
		return "", err
	}
	return string(body), nil
}

func loadJsonList(title string) (string, error) {
	filename := "data/"+title + "s.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
    }
	return string(body), nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string){
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
        port = "80"
        }
    http.HandleFunc("/car/", carHandler)
    http.HandleFunc("/comment/", commentHandler)
    http.HandleFunc("/user/", userHandler)
    http.HandleFunc("/post/", postHandler)

    //http.HandleFunc("/", frontHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
