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
	slc = append(slc, 16, 17, 18, 19, 20)

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

func e5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func e6() {
	ary := make([]string, 0, 50)
	fmt.Println("length", len(ary), "capacity", cap(ary))

	states := []string{`Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`}
	ary = append(ary, states...)
	fmt.Println("the full array:", ary)
	fmt.Println("length", len(ary), "capacity", cap(ary))

	for i, v := range ary {
		fmt.Println(i, v)
	}

	for i := 0; i < len(ary); i++ {
		fmt.Println(ary[i])
	}
}
func e7() {

	twoDstr := make([][]string, 0)
	data := []string{"James", "Bond", "Shaken, not stirred", "Miss", "Moneypenny", "Helloooooo, James"}
	for i := 0; i < 2; i++ {
		slc := []string{}
		for j := 0; j < 3; j++ {
			slc = append(slc, data[i*3+j])
		}
		twoDstr = append(twoDstr, slc)

	}
	fmt.Println(twoDstr)
}