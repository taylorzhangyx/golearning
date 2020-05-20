package main

import (
	"bufio"
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

		go scan(connection)
	}
}

func scan(connection net.Conn) {
	scner := bufio.NewScanner(connection)
	for scner.Scan() {
		line := scner.Text()
		log.Println(line)
	}
	io.WriteString(connection, "DONE!")
	defer connection.Close()
}
