package main

import (
	"fmt"
)

func main() {
	// for loops
	// print out the matrix of x,y and the values
	x, y := 10, 8
	drawMatrix(x, y)
}

func drawMatrix(x int, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("[%v,%v]\t", i, j)
		}
		fmt.Printf("\n")
	}
}
