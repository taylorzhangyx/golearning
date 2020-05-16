package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	fmt.Println("Hello, playground")
	passDataThroughChannel()
	unidirectionalChannel()
}

func passDataThroughChannel() {
	var wg sync.WaitGroup
	var c = make(chan int)

	for i := 0; i < 99; i++ {
		wg.Add(1)
		//func() {
		//}()
		go func() {
			fmt.Println(<-c)
			wg.Done()
		}()

		c <- i
	}
	wg.Wait()
}

func unidirectionalChannel() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int)
	go readChannel(c, &wg)
	go sendChannel(c)
	wg.Wait()
}

func readChannel(c <-chan int, wg *sync.WaitGroup) {
	fmt.Println(<-c)
	(*wg).Done()
}

func sendChannel(c chan<- int) {
	c <- rand.Int()
}
