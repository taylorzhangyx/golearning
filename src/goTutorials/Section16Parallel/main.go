package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	gosched()
	waitgroup()
}

func waitgroup() {
	var wg sync.WaitGroup
	exe := func(n int) {
		notuse := 0

		fmt.Println("Start", n)
		for i := 0; i < 9999; i++ {
			notuse += i
		}
		fmt.Println("Done", n)
		wg.Done()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go exe(i)
	}
	wg.Wait()
	fmt.Printf("DONE")
}

// How to use Gosched
func gosched() {
	fmt.Println("Hello, playground")
	for i := 0; i < 999; i++ {
		go func() {
			fmt.Println("fooooooooo")
		}() // goroutine
		func() {
			runtime.Gosched() // this line will tell the main thread to switch to the other go routine to process foo
			fmt.Println("bar")
		}() // this executed on the main thread and controlled by main
	}
}
