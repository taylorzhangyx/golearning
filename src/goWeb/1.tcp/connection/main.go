package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	log.Println("Start listening", listener.Addr())
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()
	count := 3
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
		}
		io.WriteString(connection, "Hello, human!")
		connection.Close()
		count--
		if count == 0 {
			log.Printf("RIP")
			break
		}
	}
}
