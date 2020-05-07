package main

import "fmt"

func main() {
	// int, string, bool, float64 are basic types in go
	// Other basic types are:
	/*
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte 	// alias for uint8

	rune 	// alias for int32
			// represents a Unicode code point

	float32 float64

	complex64 complex128

	false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
	*/
	var value = 123; // default is 0

	fmt.Println(value)
	fmt.Printf("The type of value is: %T\n", value)

	var vS = "test string"; // default is empty string

	fmt.Println(vS)
	fmt.Printf("The type of value is: %T\n", vS)

	var vF = 123.22;

	fmt.Println(vF)
	fmt.Printf("The type of value is: %T\n", vF)

	var vB = true // default is false

	fmt.Println(vB)
	fmt.Printf("The type of value is: %T\n", vB)

	s := fmt.Sprintf("a string should be printed out %s", "string")
	fmt.Println(s)
	fmt.Printf("%T", s)

}
