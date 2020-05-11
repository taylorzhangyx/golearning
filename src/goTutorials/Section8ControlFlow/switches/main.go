package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	for i := 0; i < 100; i++ {
		printSwitch(i)
	}
}

func printSwitch(v int) {
	switch v {
	case 0:
		fmt.Println(v, "Initial state")
	case 2:
		fmt.Println(v, "First even number")
	case 3, 5, 7:
		fmt.Println(v, "Prime numbers")
	case 99:
		fmt.Println(v, "last odd number")
		fallthrough
	case 101:
		fmt.Println(v, "fallthough")
	default:
		fmt.Println(v, "Nothing matches")
	}
}
