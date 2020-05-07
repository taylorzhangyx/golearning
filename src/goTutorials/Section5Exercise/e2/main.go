package main

import "fmt"

// package level declaration
// they will be assigned to zero values if no initialization values are provided.
var x int
var y string
var z bool

func main(){
	fmt.Println(x) //0
	fmt.Println(y) //""
	fmt.Println(z) //false
}
