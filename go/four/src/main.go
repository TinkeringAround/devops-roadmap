package main

import (
	"fmt"
	"log"
)

func main() {
	var err error
	var numberOfJokes = 1

	log.Println("How many Jokes you want to fetch?")
	_, err = fmt.Scan(&numberOfJokes)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("Jokes:")
	jokes, err := GetJokes(numberOfJokes)
	for _, joke := range jokes {
		log.Println("--", joke.Fact)
	}
}
