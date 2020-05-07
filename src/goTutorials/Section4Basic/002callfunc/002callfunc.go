package main

import "fmt"

func main2() {
	fmt.Println("the first line here")
	n := foo()

	fmt.Println("the print out puts", n, "letters")
}

func foo() int {
	n, _ := fmt.Println("the second line here") //  the “_” underscore character to throw away returns
	return n
}

/**
Println signature:
func Println(a ...interface{}) (n int, err error)

...<type> -  a variadic parameter
type “interface{}” is the empty interface, every value is of type “interface{}”
*/
