package api

import (
	"fmt"
	"net/http"
)

//PostCar create car object
func PostCar(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("car")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", p)
}

//GetCar get car object
func GetCar(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("car")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//PutCar updates car object
func PutCar(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("car")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//DeleteCar remove car object
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("cars")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}

//GetCarList gets car list
func GetCarList(w http.ResponseWriter, r *http.Request) {
	p, err := LoadJSON("cars")
	if err != nil {
		http.NotFound(w, r)
		fmt.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", p)
}
