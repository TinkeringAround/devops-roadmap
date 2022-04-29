package main

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	wordMap := WordCount("I am learning Go!")
	expectedWords := []string{"I", "am", "learning", "Go!"}

	for _, expectedWord := range expectedWords {
		count, ok := wordMap[expectedWord]
		if !ok {
			t.Errorf("Expected word %s was not found", expectedWord)
		}

		if count != 1 {
			t.Errorf("The count of %s was not the expected %d", expectedWord, count)
		}
	}
}
