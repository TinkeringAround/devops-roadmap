package main

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		// time.Sleep(1 * time.Second)
		_, _ = fmt.Fprintln(out, i)
	}
	_, _ = fmt.Fprint(out, "Go!")
}
