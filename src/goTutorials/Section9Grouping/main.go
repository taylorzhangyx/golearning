package main

import "fmt"

func main() {
	var ary [40]int
	for i := 0; i <= 100; i++ {
		ary[i%40] = i
	}
	fmt.Println(ary)

	fmt.Printf("type of ary is: %T", ary) // the length is included as a part of type
}
