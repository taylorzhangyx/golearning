package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
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

		go scan(connection)
	}
}

func scan(connection net.Conn) {
	// After the duration, the connection will terminate.
	connection.SetDeadline(time.Now().Add(time.Second * 10))
	scner := bufio.NewScanner(connection)
	for scner.Scan() {
		line := scner.Text()
		log.Println(line)
		io.WriteString(connection, fmt.Sprintln("Yo!", line))
	}
	defer log.Println("TIME OUT!")
	connection.Close()
}
