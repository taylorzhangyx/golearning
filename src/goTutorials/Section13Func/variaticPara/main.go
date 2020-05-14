package main

import (
	"fmt"
)

func main() {

	fmt.Println("start")

	r := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("sum is", r)
	fmt.Printf("Type is %T", r)

}

func sum(n ...int) int {
	r := 0
	for _, v := range n {
		r += v
	}
	return r
}
