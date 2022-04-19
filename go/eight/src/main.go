package main

import "fmt"

func fibonacci() func() int {
	var a, b = -1, 1

	return func() int {
		var res = a + b
		a = b
		b = res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
