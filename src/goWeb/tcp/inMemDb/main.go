package main

import (
	"bufio"
	"fmt"
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

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go memDb(connection)
	}
}

func memDb(connection net.Conn) {
	scner := bufio.NewScanner(connection)
	for scner.Scan() {
		line := scner.Text()
		log.Println(line)
		io.WriteString(connection, fmt.Sprintln("Yo!", line))
	}
	defer log.Println("TIME OUT!")
	connection.Close()
}
