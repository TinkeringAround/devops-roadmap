// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import (
	"fmt"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	b[0] = 65 // 65 = "A"
	return 1, nil
}

func Readers() {
	fmt.Println("## Readers Exercise ##")
	r := MyReader{}

	b := make([]byte, 8)
	for i := 0; i < 1; i++ {
		n, err := r.Read(b)
		fmt.Printf("n = %x, err = %v, b = %v, b[:n] = %q\n", n, err, b, b[:n])
	}
}
