package main

import "fmt"

type rank int

func main(){
	var x int = 10
	var y rank = 9
	fmt.Println(x)
	fmt.Printf("type of x is: %T", x)
	fmt.Println(y)
	fmt.Printf("type of y is: %T", y)

	// convert one type to another if compatable
	y = rank(x)
	fmt.Println(y)
}
