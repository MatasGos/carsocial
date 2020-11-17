package api

import (
	"fmt"
	"net/http"
)

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
