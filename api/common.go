package api

import (
	"database/sql"
	"io/ioutil"
)

const (
	host     = "database-1.carhd0fhradi.eu-central-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "astonmartin"
	dbname   = "carsocial"
)

//PsqlInfo database info
const PsqlInfo = "host=database-1.carhd0fhradi.eu-central-1.rds.amazonaws.com " +
	"port=5432 " +
	"user=postgres " +
	"password=astonmartin " +
	"dbname=carsocial " +
	"sslmode=disable "

//LoadJSON returns json from text file
func LoadJSON(title string) (string, error) {
	filename := "data/" + title + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func loadJSONList(title string) (string, error) {
	filename := "data/" + title + "s.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//Database is database connection
var Database *sql.DB

//OpenDatabase opens connection to database
func OpenDatabase() {
	var err error
	Database, err = sql.Open("postgres", PsqlInfo)

	err = Database.Ping()
	if err != nil {
		panic(err)
	}
}

//CloseDatabase closes connection to database
func CloseDatabase() {
	Database.Close()
}
