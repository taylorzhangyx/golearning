package main

import (
	"bufio"
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
		scanner := bufio.NewScanner(con)
		go func(con net.Conn) {
			for scanner.Scan() {
				str := scanner.Text()
				_, err := io.WriteString(con, fmt.Sprintln("I see you connected:", str))
				if err != nil {
					panic(fmt.Sprintln("Write Failed", err))
				}
				log.Println(str)
				if str == "" {
					break
				}
			}
			count++
			defer con.Close()
		}(con)
		if count >= 5 {
			lis.Close()
			break
		}
	}
	defer log.Println("Server RIP")
}
