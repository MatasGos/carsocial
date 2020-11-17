package api

import (
	"io/ioutil"
)

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
