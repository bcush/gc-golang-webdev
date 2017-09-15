package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		// Request
		request(conn)
	}
}

func request(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn)
}
