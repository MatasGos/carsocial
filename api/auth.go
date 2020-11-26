package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Token is
type Token struct {
	Token string `json:"token"`
}

var TokenAuth *jwtauth.JWTAuth

// func init() {
// 	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

// 	// For debugging/example purposes, we generate and print
// 	// a sample jwt token with claims `user_id:123` here:
// 	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
// 	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
// }

//Login good
func Login(w http.ResponseWriter, r *http.Request) {
	var temp login
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	sqlQ := "SELECT password, role, id " +
		"FROM public.users WHERE email=$1"

	row := Database.QueryRow(sqlQ, temp.Email)
	var user User
	err = row.Scan(&user.Password, &user.Role, &user.ID)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested user doesnt exist", http.StatusNotFound)
		return
	case nil:
	default:
		panic(err)
	}

	if user.Password == temp.Password {
		token := createToken(user.ID, user.Role)
		var tokenObj Token
		tokenObj.Token = token
		json, err := json.Marshal(tokenObj)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json)

	} else {
		http.Error(w, "bad login", http.StatusBadRequest)
	}
}

func createToken(id int, role string) string {

	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 60).Unix()})
	fmt.Println(id)
	fmt.Println(role)
	return tokenString
}

//
//localhost:5000/cars/
