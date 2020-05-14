package main

import "fmt"

func main() {
	// fib := initFib()
	// for i := 0; i < 10; i++ {
	// 	v := fib() // return the first 10 fib values
	// 	fmt.Println(i, v)
	// }

	recurFib := reverseFib(20)
	recurFib()
}

// closure way to print out the fib numbers
func initFib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func reverseFib(count int) func() {
	a, b := 1, 1
	return func() {
		a, b = fib(a, b, count)
	}
}

func fib(a int, b int, n int) (int, int) {
	a, b = b, a+b
	if n != 0 {
		n--
		fmt.Println(n, a)
		fib(a, b, n)
	} else {
		fmt.Println("DONE")
	}
	return a, b
}
