/*
Let's have some fun with functions.
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

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

func Closures() {
	fmt.Println("## Closures Exercise ##")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
