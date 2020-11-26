package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

//Post is post
type Post struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Fkcar  int    `json:"fk_car"`
	Fkuser int    `json:"fk_user"`
}

type postSQL struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Fkcar  int    `json:"fk_car"`
	Fkuser int    `json:"fk_user"`
}

//PostPost create post object
func PostPost(w http.ResponseWriter, r *http.Request) {
	var post postSQL
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	sql := "INSERT INTO public.posts(" +
		"text, fk_car, fk_user)" +
		"VALUES ($1, $2, $3);"

	err = Database.QueryRow(sql, post.Text, post.Fkcar, post.Fkuser).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	GetPostList(w, r)
	w.WriteHeader(http.StatusCreated)
}

//GetPost gets post object
func GetPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	sqlQ := "SELECT 	id  ," +
		"text ," +
		"fk_car, fk_user FROM public.posts WHERE id=$1"

	row := Database.QueryRow(sqlQ, postID)

	var post Post
	err := row.Scan(&post.ID, &post.Text, &post.Fkcar, &post.Fkuser)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested post no longer exists", http.StatusNotFound)
		return
	case nil:
		if err != nil {
			panic(err)
		}

		json, err := json.Marshal(post)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json)
	default:
		panic(err)
	}
}

//PutPost updates post object
func PutPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	w.Header().Set("Content-Type", "application/json")

	//read body
	var newData Post
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &newData)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}
	//update
	sqlPut := "UPDATE public.posts SET"
	if newData.Text != "" {
		sqlPut += " Text='" + fmt.Sprint(newData.Text) + "'"
	}
	sqlPut += " WHERE id=" + fmt.Sprint(postID) + ";"

	err = Database.QueryRow(sqlPut).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}
	fmt.Println(sqlPut)
}

//DeletePost delets post object
func DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	sql := "DELETE FROM public.posts WHERE id=$1;"

	err := Database.QueryRow(sql, postID).Err()
	if err != nil {
		http.Error(w, "Failed delete", http.StatusBadRequest)
		panic(err)
	}

	GetPostList(w, r)
}

//GetPostList gets posts list
func GetPostList(w http.ResponseWriter, r *http.Request) {
	sqlQ := "SELECT 	id  ," +
		"text ," +
		"fk_car, fk_user FROM public.posts"

	rows, err := Database.Query(sqlQ)
	if err != nil {
		panic(err)
	}
	var posts [20]Post
	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&posts[count].ID, &posts[count].Text, &posts[count].Fkcar, &posts[count].Fkuser)
		if err != nil {
			panic(err)
		}
		count++
	}
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(posts[:count])

	fmt.Fprintf(w, "%s", json)
}
