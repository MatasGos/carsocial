package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

//Car is car
type Car struct {
	ID           int       `json:"id"`
	Model        string    `json:"model"`
	Manufacturer string    `json:"manufacturer"`
	Plate        string    `json:"plate"`
	Color        string    `json:"color"`
	Caradded     time.Time `json:"caradded"`
	Year         string    `json:"year"`
	Fkuser       string    `json:"fkuser"`
	Vin          string    `json:"vin"`
}

type carSQL struct {
	ID           int
	Model        string
	Manufacturer string
	Plate        sql.NullString
	Color        sql.NullString
	Caradded     sql.NullTime
	Year         sql.NullString
	Fkuser       int
	Vin          sql.NullString
}

//PostCar create car object
func PostCar(w http.ResponseWriter, r *http.Request) {

	var cars []Car
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &cars)
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	for i := range cars {
		sql := "INSERT INTO public.cars(" +
			"model, manufacturer, plate, color, caradded, year, fk_user, vin)" +
			"VALUES ($1, $2, $3, $4, CURRENT_DATE, $5, $6, $7);"

		err = Database.QueryRow(sql, cars[i].Model, cars[i].Manufacturer, cars[i].Plate, cars[i].Color, cars[i].Year, cars[i].Fkuser, cars[i].Vin).Err()
		if err != nil {
			http.Error(w, "wrong body structure", http.StatusBadRequest)
			panic(err)
		}
	}

	GetCarList(w, r)
	w.WriteHeader(http.StatusCreated)
}

//GetCar get car object
func GetCar(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "carID")
	sqlQ := "SELECT 	id  ," +
		"model ," +
		"manufacturer, " +
		"plate,  " +
		"color, " +
		"caradded, " +
		"year,  " +
		"fk_user, vin FROM cars WHERE id=$1"

	row := Database.QueryRow(sqlQ, carID)

	var cars Car
	var car carSQL
	err := row.Scan(&cars.ID, &car.Model, &cars.Manufacturer, &car.Plate,
		&car.Color, &car.Caradded, &car.Year, &cars.Fkuser, &car.Vin)

	switch err {
	case sql.ErrNoRows:
		http.Error(w, "requested car no longer exists", http.StatusNotFound)
		return
	case nil:
		if err != nil {
			panic(err)
		}
		if car.Plate.Valid {
			cars.Plate = string(car.Plate.String)
		}
		if car.Color.Valid {
			cars.Color = string(car.Color.String)
		}
		if car.Caradded.Valid {
			cars.Caradded = time.Time(car.Caradded.Time)
		}
		if car.Year.Valid {
			cars.Year = string(car.Year.String)
		}
		if car.Vin.Valid {
			cars.Vin = string(car.Vin.String)
		}

		json, err := json.Marshal(cars)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json)
	default:
		panic(err)
	}

}

//PutCar updates car object
func PutCar(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "carID")
	w.Header().Set("Content-Type", "application/json")

	//read body
	var newData Car
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

	fmt.Println(json.Marshal(newData.Model))
	fmt.Println(json.Marshal(newData.Manufacturer))
	if newData.Plate == "" {
		fmt.Println("juozasdx")
	}
	fmt.Println(json.Marshal(newData.Vin))
	//update
	sqlPut := "UPDATE public.cars SET"
	if newData.Model != "" {
		sqlPut += " model='" + fmt.Sprint(newData.Model) + "'"
	}
	if newData.Manufacturer != "" {
		sqlPut += " manufacturer='" + fmt.Sprint(newData.Manufacturer) + "'"
	}
	if newData.Plate != "" {
		sqlPut += " late='" + fmt.Sprint(newData.Plate) + "'"
	}
	if newData.Color != "" {
		sqlPut += " color='" + fmt.Sprint(newData.Color) + "'"
	}
	if newData.Year != "" {
		sqlPut += " year='" + fmt.Sprint(newData.Year) + "'"
	}
	if newData.Vin != "" {
		sqlPut += " vin='" + fmt.Sprint(newData.Vin) + "'"
	}
	sqlPut += " WHERE id=" + fmt.Sprint(carID) + ";"

	err = Database.QueryRow(sqlPut).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}
	fmt.Println(sqlPut)

}

//DeleteCar remove car object
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "carID")
	fmt.Println(carID)

	sql := "DELETE FROM public.cars WHERE id=$1;"

	err := Database.QueryRow(sql, carID).Err()
	if err != nil {
		http.Error(w, "wrong body structure", http.StatusBadRequest)
		panic(err)
	}

	GetCarList(w, r)
}

//GetCarList gets car list
func GetCarList(w http.ResponseWriter, r *http.Request) {

	sql := "SELECT 	id  ," +
		"model ," +
		"manufacturer, " +
		"plate,  " +
		"color, " +
		"caradded, " +
		"year,  " +
		"fk_user, vin FROM cars"

	rows, err := Database.Query(sql)
	var cars [20]Car
	count := 0
	defer rows.Close()
	for rows.Next() {
		var car carSQL
		err = rows.Scan(&cars[count].ID, &cars[count].Model, &cars[count].Manufacturer, &car.Plate, &car.Color, &car.Caradded, &car.Year, &cars[count].Fkuser, &car.Vin)
		if err != nil {
			panic(err)
		}
		if car.Plate.Valid {
			cars[count].Plate = string(car.Plate.String)
		}
		if car.Color.Valid {
			cars[count].Color = string(car.Color.String)
		}
		if car.Caradded.Valid {
			cars[count].Caradded = time.Time(car.Caradded.Time)
		}
		if car.Year.Valid {
			cars[count].Year = string(car.Year.String)
		}
		if car.Vin.Valid {
			cars[count].Vin = string(car.Vin.String)
		}
		count++
	}

	json, err := json.Marshal(cars[:count])

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)
}
