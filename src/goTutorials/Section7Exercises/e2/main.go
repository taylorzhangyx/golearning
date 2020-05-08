package main

import (
	"fmt"
)

//Create the typed and untyped constants, print out the values

func main() {

	const (
		a      = 33
		ta int = 44
	)
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", ta)
	fmt.Printf("%T\n", ta)
}
