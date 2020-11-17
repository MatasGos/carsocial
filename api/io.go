package api

import (
	"io/ioutil"
)

func LoadJSON(title string) (string, error) {
	filename := "data/" + title + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func LoadJSONList(title string) (string, error) {
	filename := "data/" + title + "s.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
