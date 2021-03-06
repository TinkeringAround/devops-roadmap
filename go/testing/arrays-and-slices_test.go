package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum all numbers in array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should sum all", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
