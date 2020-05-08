package main

import (
	"fmt"
	"runtime"
)

const (
	a = iota
	b
	c
	d
	e
)

const (
	bitshift1 = 1 << iota
	bitshift2 = 1 << iota
	bitshift3 = 1 << iota
	bitshift4 = 1 << iota
	bitshift5 = 1 << iota
)

func main() {
	// get the computer information
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println("iota: ", a)
	fmt.Println("iota: ", b)
	fmt.Println("iota: ", c)
	fmt.Println("iota: ", d)
	fmt.Println("iota: ", e)

	fmt.Println("bitshiftiota: ", bitshift1)
	fmt.Println("bitshiftiota: ", bitshift2)
	fmt.Println("bitshiftiota: ", bitshift3)
	fmt.Println("bitshiftiota: ", bitshift4)
	fmt.Println("bitshiftiota: ", bitshift5)

	unsignedConst()
}

func unsignedConst() {
	const (
		cint        = 10 // untyped const
		cflo        = 4.3
		cstr        = "str"
		ctint   int = 10 // typed const
		cbigint     = 99999
	)
	type myInt int
	var testInt myInt = cint // valid
	fmt.Println(testInt)

	//testInt = ctint // error - type does not match
	//var smlint int8 = cbigint // error - the compiler warns a overflow
	// since the compiler knows the value of the constant, it will evaluate the assignment
}
