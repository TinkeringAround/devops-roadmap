package main

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(2, 7)
	fmt.Println(sum)
	// Output: 9
}

func TestAdd(t *testing.T) {
	t.Run("should add two numbers", func(t *testing.T) {
		sum := Add(2, 4)
		want := 6

		if sum != want {
			t.Errorf("expected '%d' but got '%d'", want, sum)
		}
	})
}
