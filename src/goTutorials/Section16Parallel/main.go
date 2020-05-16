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

func racecondition() {
	var wg sync.WaitGroup
	exe := func(n int) {
		notuse := 0
		for i := 0; i < 9999; i++ {
			notuse += i
		}
		fmt.Println("Done", n)
		wg.Done()
	}

	count := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			counter := count

			runtime.Gosched()
			go exe(i)
			count = counter + 1
		}(i)
	}
	wg.Wait()
	fmt.Printf("DONE, %v", count)
}

func mutex() {
	var m sync.Mutex
	var wg sync.WaitGroup
	exe := func(n int) {
		notuse := 0
		for i := 0; i < 9999; i++ {
			notuse += i
		}
		fmt.Println("Done", n)
		wg.Done()
	}

	count := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Lock()
			counter := count

			runtime.Gosched()
			go exe(i)
			count = counter + 1
			m.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("DONE, %v", count)
}
