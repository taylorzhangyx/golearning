package main

import "fmt"

// package level declaration
// they will be assigned to zero values if no initialization values are provided.
var x int
var y string
var z bool

func main(){
	s:= fmt.Sprintf("x:%d, y:%s, z:%t",x,y,z)
	fmt.Println(s)
}
