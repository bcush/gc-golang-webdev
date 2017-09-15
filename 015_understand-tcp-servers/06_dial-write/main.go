package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8089")
	if err != nil {
		log.Fatalln("Your err is", err)
	}

	fmt.Fprintln(conn, "Hello, I dialed you!")
}
