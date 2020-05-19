// Something to document
package main

import (
	"fmt"
)

// type error interface {
// 	Error() string
// }

// https://blog.golang.org/defer-panic-and-recover
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil { // Revocer is only useful in deferred function
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	G(0)
	fmt.Println("Returned normally from g.") // This line will not be executed since the panic popped up the function
}

// G Document for G.
func G(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	G(i + 1)
}
