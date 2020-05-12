package main

import "fmt"

func main() {
	defineAry()
	makeaSlice()
}

func defineAry() {
	var ary [40]int
	for i := 0; i <= 100; i++ {
		ary[i%40] = i
	}
	fmt.Println(ary)

	fmt.Printf("type of ary is: %T", ary) // the length is included as a part of type
}

func makeaSlice() {

	slc := make([]int, 0)
	fmt.Println(slc)

	for i := 0; i < 10; i++ {
		slc = append(slc, i)
	}
	fmt.Println(slc)

	for i, v := range slc {
		fmt.Println(i, v)
	}

}
