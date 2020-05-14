package main

import "fmt"

func main() {
	var ary [5]int
	ary[0] = 10
	ary[1] = 11
	ary[2] = 12
	ary[3] = 13
	ary[4] = 14

	for i, v := range ary {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", ary)
}
