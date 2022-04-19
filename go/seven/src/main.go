package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wordMap[word]++
	}
	return wordMap
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
}
