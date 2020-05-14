package main

import (
	"fmt"
)

func main() {
	fib := initFib()
	for i := 0; i < 10; i++ {
		v := fib() // return the first 10 fib values
		fmt.Println(i, v)
	}
}

func initFib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
