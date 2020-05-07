package main

import "fmt"

type taylor int

func main() {
	var x taylor
	fmt.Println(x)
	fmt.Printf("type of x: %T\n", x)
	x = 42
	fmt.Println(x)

	y := int(x)
	fmt.Println(y)
}
