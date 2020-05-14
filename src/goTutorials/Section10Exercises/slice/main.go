package main

import "fmt"

func main() {
	e3()
	e4()
}

func e3() {
	var slc []int
	//or slc := []int{}
	//or slc := make([]int, 0, 100)
	slc = append(slc, 11)
	slc = append(slc, 12)
	slc = append(slc, 13)
	slc = append(slc, 14)
	slc = append(slc, 15)
	slc = append(slc, 16)
	slc = append(slc, 17)
	slc = append(slc, 18)
	slc = append(slc, 19)
	slc = append(slc, 20)

	for i, v := range slc {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slc)
}

func e4() {

	ary := make([]int, 0, 20)

	y := []int{25, 26, 27, 28}
	ary = append(ary, 21)
	ary = append(ary, 22)
	ary = append(ary, 23)
	ary = append(ary, 23)
	ary = append(ary, y...)

	for i, v := range ary {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", ary)

}
