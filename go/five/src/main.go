package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Println(strconv.Itoa(i)+".", "Iteration:", sqrt(2, i))
	}
}

func sqrt(x float64, iterations int) float64 {
	z := 1.0
	for i := 0; i < iterations; i++ {
		z -= (z*z - x) / (x * z)
	}

	return z
}
