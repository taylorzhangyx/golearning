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

func rangeChannel() {
	// range will check the channel untill it's closed
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- rand.Intn(99)
		}
		close(c)
	}()
	for v := range c { // it'll pull all the remaining values in the channel after the channel is closed
		fmt.Println(v)
	}
}
