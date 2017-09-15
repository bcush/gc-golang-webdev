package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "Hello from WriteString")
		fmt.Fprintln(conn, "Hello from Fprintln!")
		fmt.Fprintf(conn, "%v", "Hello from Fprintf!")

		conn.Close()
	}
}
