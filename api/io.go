package api

import (
	"io/ioutil"
)

func loadJson(title string) (string, error) {
	filename := "data/" + title + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func loadJsonList(title string) (string, error) {
	filename := "data/" + title + "s.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
