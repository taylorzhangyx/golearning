package main

import (
	"fmt"
)

func main() {
	// for loops
	// print out the matrix of x,y and the values
	x, y := 10, 8
	drawMatrix(x, y)

	// print 5
	max := 243
	print5(max)
}

func drawMatrix(x int, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("[%v,%v]\t", i, j)
		}
		fmt.Printf("\n")
	}
}

func print5(rng int) {
	for i := 0; i < rng; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}
}

func printChars() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%c", i)
	}
}
