package main

import "fmt"

type taylor int

func main() {
	var x taylor
	fmt.Println(x)
	fmt.Printf("type of x: %T", x)
	x = 42
	fmt.Println(x)
}
