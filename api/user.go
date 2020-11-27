package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

//User is user
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	UserCreated time.Time `json:"usercreated"`
	Role        string    `json:"role"`
	Password    string    `json:"password"`
}

type userSQL struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Username    string         `json:"username"`
	Phone       sql.NullString `json:"phone"`
	Email       string         `json:"email"`
	UserCreated time.Time      `json:"usercreated"`
	Role        string         `json:"role"`
	Password    string         `json:"password"`
}

//PostUser create user object
func PostUser(w http.ResponseWriter, r *http.Request) {
	var users []User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	for i := range users {
		sql := "INSERT INTO public.users(" +
			"name, username, phone, email, usercreated, role, password)" +
			"VALUES ($1, $2, $3, $4, CURRENT_DATE, $5, $6);"

		err = Database.QueryRow(sql, users[i].Name, users[i].Username, users[i].Phone, users[i].Email, users[i].Role, users[i].Password).Err()
		if err != nil {
			http.Error(w, "wrong data", http.StatusBadRequest)
			panic(err)
		}
	}

	GetUserList(w, r)
	w.WriteHeader(http.StatusCreated)
}

//GetUser gets user object
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	sqlQ := "SELECT 	id, name, username, phone, email, usercreated, role, password " +
		"FROM public.users WHERE id=$1"

	row := Database.QueryRow(sqlQ, userID)

	var user User
	var userSQL userSQL
	err := row.Scan(&user.ID, &user.Name, &user.Username, &userSQL.Phone,
		&user.Email, &user.UserCreated, &user.Role, &user.Password)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested user no longer exists", http.StatusNotFound)
		return
	case nil:
		if err != nil {
			panic(err)
		}
		if userSQL.Phone.Valid {
			user.Phone = string(userSQL.Phone.String)
		}

		json, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json)
	default:
		panic(err)
	}
}

//PutUser updates user object
func PutUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	sqlQ := "SELECT 	id " +
		"FROM public.users WHERE id=$1"

	row := Database.QueryRow(sqlQ, userID)

	var user User
	err := row.Scan(&user.ID)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested user no longer exists", http.StatusNotFound)
		return
	case nil:
	default:
		panic(err)
	}
	token, err := TokenAuth.Decode(jwtauth.TokenFromHeader(r))
	if token.Claims.(jwt.MapClaims)["id"] != user.ID && token.Claims.(jwt.MapClaims)["role"] != "admin" {
		http.Error(w, "Unauthorized action", http.StatusUnauthorized)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	//read body
	var newData User
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
	count := 0
	sqlPut := "UPDATE public.users SET"
	if newData.Name != "" {
		sqlPut += " name='" + fmt.Sprint(newData.Name) + "'"
		count++
	}
	if newData.Username != "" {
		if count > 0 {
			sqlPut += ","
		}
		count++
		sqlPut += " username='" + fmt.Sprint(newData.Username) + "'"
	}
	if newData.Phone != "" {
		if count > 0 {
			sqlPut += ","
		}
		count++
		sqlPut += " phone='" + fmt.Sprint(newData.Phone) + "'"
	}
	if newData.Email != "" {
		if count > 0 {
			sqlPut += ","
		}
		count++
		sqlPut += " email='" + fmt.Sprint(newData.Email) + "'"
	}
	if newData.Password != "" {
		if count > 0 {
			sqlPut += ","
		}
		count++
		sqlPut += " password='" + fmt.Sprint(newData.Password) + "'"
	}
	if newData.Role != "" {
		if count > 0 {
			sqlPut += ","
		}
		count++
		sqlPut += " role='" + fmt.Sprint(newData.Role) + "'"
	}
	sqlPut += " WHERE id=" + fmt.Sprint(userID) + ";"

	err = Database.QueryRow(sqlPut).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}
}

//DeleteUser deletes user object
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	sqlQ := "SELECT 	id " +
		"FROM public.users WHERE id=$1"

	row := Database.QueryRow(sqlQ, userID)

	var user User
	err := row.Scan(&user.ID)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested user no longer exists", http.StatusNotFound)
		return
	case nil:
	default:
		panic(err)
	}
	token, err := TokenAuth.Decode(jwtauth.TokenFromHeader(r))
	if token.Claims.(jwt.MapClaims)["id"] != user.ID && token.Claims.(jwt.MapClaims)["role"] != "admin" {
		http.Error(w, "Unauthorized action", http.StatusUnauthorized)
		panic(err)
	}
	sql := "DELETE FROM public.users WHERE id=$1;"

	err = Database.QueryRow(sql, userID).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	GetUserList(w, r)
}

//GetUserList returns users list
func GetUserList(w http.ResponseWriter, r *http.Request) {
	sqlQ := "SELECT 	id, name, username, phone, email, usercreated, role, password FROM users"

	rows, err := Database.Query(sqlQ)
	if err != nil {
		panic(err)
	}

	var users [20]User
	count := 0
	defer rows.Close()
	for rows.Next() {
		var userSQL userSQL
		err := rows.Scan(&users[count].ID, &users[count].Name, &users[count].Username, &userSQL.Phone,
			&users[count].Email, &users[count].UserCreated, &users[count].Role, &users[count].Password)
		if err != nil {
			panic(err)
		}
		if userSQL.Phone.Valid {
			users[count].Phone = string(userSQL.Phone.String)
		}
		count++
	}
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(users[:count])
	fmt.Fprintf(w, "%s", json)
}
