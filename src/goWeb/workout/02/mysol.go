package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	log.Println("Start App")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(fmt.Sprintln("Listen Failed", err))
	}

	count := 0
	for {
		con, err := lis.Accept()
		if err != nil {
			panic(fmt.Sprintln("Accept Failed", err))
		}
		go func(con net.Conn) {
			_, err := io.WriteString(con, "I see you connected")
			if err != nil {
				panic(fmt.Sprintln("Write Failed", err))
			}
			count++
			defer con.Close()
		}(con)
		if count >= 1 {
			lis.Close()
			break
		}
	}
	defer log.Println("Server RIP")
}
