package main // the entry point of a module, or a file. the package main is required to execute a module.
// otherwise, go run: cannot run non-main package wil be displayed

import "fmt" //importing the package you use in this file

func main1() { // func main is also another requirement to execute the package.
	// otherwise: runtime.main_main·f: function main is undeclared in the main package is displayed
	fmt.Println("hello world")
}
