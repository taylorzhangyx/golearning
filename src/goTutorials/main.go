package main

import (
	"fmt"
	"runtime"
)

func main() {
	// get the computer information
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
