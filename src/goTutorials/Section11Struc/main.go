package main

import (
	"fmt"
)

func main() {
	type person struct {
		first   string
		last    string
		birth   int
		favIceC string
	}

	p1 := person{
		first:   "Yuxin",
		last:    "Zhang",
		birth:   1992,
		favIceC: "Choc",
	}

	p2 := struct {
		count  int
		person []person
	}{
		count:  2,
		person: []person{p1, {first: "Taylor", last: "Zhang", birth: 22}},
	}

	fmt.Println(p1, p2)

}
