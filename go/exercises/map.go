/* Implement WordCount.
It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.
You might find strings.Fields helpful.
*/

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

func Maps() {
	fmt.Println("## Map Exercise ##")
	fmt.Println(WordCount("I am learning Go!"))
}
