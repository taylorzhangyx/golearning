package main

import "fmt"

func main() {

	var i num = 12
	i.repeat()
}

type num int

func (s num) repeat() {
	fmt.Println(s, s)
}
