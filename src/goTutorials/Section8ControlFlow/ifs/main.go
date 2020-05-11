package main

import "fmt"

func main() {
	fmt.Println("Start")
	// Narrow the scope of the veriable by init it in if
	param := "name"

	if name := param; name == "name" {
		// variable name is only valid inside if
		fmt.Println("the value of name is name")
	}
}
