/*Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
(Use uint8(intValue) to convert between types.)*/

package main

import (
	"fmt"
)

func generatePicture(dx, dy int) [][]uint8 {
	var picture [][]uint8
	for x := 0; x < dx; x++ {
		var row []uint8
		for y := 0; y < dy; y++ {
			row = append(row, uint8(x*y))
		}

		picture = append(picture, row)
	}

	return picture
}

func Slice() {
	fmt.Println("## Picture (Slice) Exercise ##")
	fmt.Println(generatePicture(8, 8))
}
