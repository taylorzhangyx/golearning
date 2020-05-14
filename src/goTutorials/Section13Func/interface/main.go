package main

import (
	"fmt"
)

// There is no inheretance in go.

func main() {
	fmt.Println("Start")
	a1 := cloths{
		sleeve: true,
		color:  "blk",
	}
	fmt.Println("a1:", a1)

	a2 := tshirt{
		cloths: cloths{
			sleeve: true,
			color:  "blk",
		},
		print: false,
	}
	fmt.Println("a2:", a2)

	a3 := suit{
		cloths: cloths{

			sleeve: true,
			color:  "blk",
		},
		tie: true,
	}
	fmt.Println("a3:", a3)
	a1.putOn()
	a2.putOn()
	a3.putOn()

	person(a3)
}

type cloths struct {
	sleeve bool
	color  string
}

type tshirt struct {
	cloths
	print bool
}

type suit struct {
	cloths
	tie bool
}

type wear interface {
	putOn()
}

func (c cloths) putOn() {
	fmt.Println("cloths -- 111")
}

func (t tshirt) putOn() {
	fmt.Println("tshirt -- 22")
}

func (s suit) putOn() {
	fmt.Println("suit -- 33")
}

func person(c wear) {
	fmt.Println("Wearing:", c)
	a, ok1 := c.(cloths)
	b, ok2 := c.(suit)
	a.putOn()
	fmt.Println("a:", a, ",ok:", ok1)
	fmt.Println("b:", b, ",ok:", ok2)
	c.(suit).putOn()
}
