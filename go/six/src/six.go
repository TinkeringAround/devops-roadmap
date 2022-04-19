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

func main() {
	fmt.Println(generatePicture(8, 8))
}
