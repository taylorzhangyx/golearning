package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	db := make(map[string]string)
	for scner.Scan() {
		line := scner.Text()
		log.Println("command:", line, connection.LocalAddr(), connection.RemoteAddr())

		command, data := split(line)

		res := work(&db, command, data)

		io.WriteString(connection, fmt.Sprintln("Db is working:", res))
	}
	connection.Close()
}

func work(db *map[string]string, cmd string, data string) string {
	var res string

	tokens := strings.Split(data, " ")
	key := tokens[0]
	val := strings.Join(tokens[1:], " ")

	switch cmd {
	case "ADD":
		(*db)[key] = val
		res = fmt.Sprintln(key, ":", val, "is saved")
	case "DEL":
		delete(*db, key)
		res = fmt.Sprintln(key, "is deleted")
	case "GET":
		res = (*db)[key]
	default:
		res = "Invalid! Try again."
	}
	return res
}

func split(line string) (string, string) {
	tokens := strings.Split(line, " ")
	data := strings.Join(tokens[1:], " ")
	return tokens[0], data
}
