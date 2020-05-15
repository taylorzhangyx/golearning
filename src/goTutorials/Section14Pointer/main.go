package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	a := 42
	b := "something"
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)

	a1 := &a
	*a1 = 33
	fmt.Println(a)

	p := person{
		first: "tay",
		last:  "lor",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}

func changeMe(p *person) {
	p.first = "looking"
}

/*******
When to use pointers
1. When load a big chunk of data, pass the pointer of the data to read them piece by piece
2. When to modify a value of the given location only

In Go, everything is passed by value.
*/
