package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

//Comment is comment
type Comment struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Fkpost   int    `json:"fk_post"`
	Fkuser   int    `json:"fk_user"`
	Username string `json:"username"`
}

type commentSQL struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Fkpost int    `json:"fk_post"`
	Fkuser int    `json:"fk_user"`
}

//PostComment creates comment object
func PostComment(w http.ResponseWriter, r *http.Request) {
	claims := GetToken(jwtauth.TokenFromHeader(r))
	var comment commentSQL
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &comment)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	sql := "INSERT INTO public.comments(" +
		"text, fk_post, fk_user)" +
		"VALUES ($1, $2, $3);"

	err = Database.QueryRow(sql, comment.Text, comment.Fkpost, claims["id"]).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	GetCommentList(w, r)
	w.WriteHeader(http.StatusCreated)
}

//GetComment returns comment object
func GetComment(w http.ResponseWriter, r *http.Request) {
	//jwauth.Verifier(r.Header.Get("Autohrization"))

	commentID := chi.URLParam(r, "commentID")
	sqlQ := "SELECT 	id  ," +
		"text ," +
		"fk_post, fk_user FROM public.comments WHERE id=$1"

	row := Database.QueryRow(sqlQ, commentID)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.Text, &comment.Fkpost, &comment.Fkuser)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested comment no longer exists", http.StatusNotFound)
		return
	case nil:
		if err != nil {
			panic(err)
		}

		json, err := json.Marshal(comment)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json)
	default:
		panic(err)
	}
}

//PutComment updates comment object
func PutComment(w http.ResponseWriter, r *http.Request) {

	commentID := chi.URLParam(r, "commentID")
	sqlQ := "SELECT 	id  ," +
		"text ," +
		"fk_post, fk_user FROM public.comments WHERE id=$1"

	row := Database.QueryRow(sqlQ, commentID)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.Text, &comment.Fkpost, &comment.Fkuser)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested comment no longer exists", http.StatusNotFound)
		return
	case nil:
	default:
		panic(err)
	}

	claims := GetToken(jwtauth.TokenFromHeader(r))
	if claims["id"] != comment.Fkuser && claims["role"] != "admin" {
		http.Error(w, "Unauthorized action", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	//read body
	var newData Comment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &newData)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
	}
	//update
	sqlPut := "UPDATE public.comments SET"
	if newData.Text != "" {
		sqlPut += " Text='" + fmt.Sprint(newData.Text) + "'"
	}
	sqlPut += " WHERE id=" + fmt.Sprint(commentID) + ";"

	err = Database.QueryRow(sqlPut).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
	}
	fmt.Println(sqlPut)

}

//DeleteComment removes comment object
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentID")
	sqlQ := "SELECT 	id  ," +
		"text ," +
		"fk_post, fk_user FROM public.comments WHERE id=$1"

	row := Database.QueryRow(sqlQ, commentID)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.Text, &comment.Fkpost, &comment.Fkuser)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested comment no longer exists", http.StatusNotFound)
		return
	case nil:
	default:
		panic(err)
	}

	claims := GetToken(jwtauth.TokenFromHeader(r))
	if claims["id"] != comment.Fkuser && claims["role"] != "admin" {
		http.Error(w, "Unauthorized action", http.StatusUnauthorized)
		return
	}

	sql := "DELETE FROM public.comments WHERE id=$1;"

	err = Database.QueryRow(sql, commentID).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		return
	}

	GetCommentList(w, r)
}

//GetCommentList returns comment list
func GetCommentList(w http.ResponseWriter, r *http.Request) {
	sqlQ := "SELECT 	id  ," +
		"text ," +
		"fk_post, fk_user FROM public.comments"

	rows, err := Database.Query(sqlQ)
	if err != nil {
		panic(err)
	}
	var comments [20]Comment
	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&comments[count].ID, &comments[count].Text, &comments[count].Fkpost, &comments[count].Fkuser)
		if err != nil {
			panic(err)
		}
		count++
	}

	json, err := json.Marshal(comments[:count])

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)

}

//GetCommentFull returns comment list
func GetCommentFull(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	sqlQ := "SELECT 	comments.id as id ," +
		"text , fk_post, fk_user username FROM public.comments left join public.users   ON public.users.id=public.comments.fk_user " +
		"WHERE fk_post=$1"

	rows, err := Database.Query(sqlQ, postID)
	if err != nil {
		panic(err)
	}
	var comments [20]Comment
	count := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&comments[count].ID, &comments[count].Text, &comments[count].Fkpost, &comments[count].Fkuser)
		if err != nil {
			panic(err)
		}
		count++
	}

	json, err := json.Marshal(comments[:count])

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)

}
