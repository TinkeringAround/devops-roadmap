package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Joke struct {
	Fact string
}

func GetJokes(amount int) (jokes []Joke, err error) {
	// Fetching
	response, err := http.Get("https://dog-facts-api.herokuapp.com/api/v1/resources/dogs?number=" + strconv.Itoa(amount))
	if err != nil {
		return
	}

	// Reading Body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// Converting Body to Array of Joke Struct
	err = json.Unmarshal(body, &jokes)
	if err != nil {
		return
	}

	return
}
