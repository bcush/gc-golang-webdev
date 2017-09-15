package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections
	li, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	// Close the listener when the application closes
	defer li.Close()
	fmt.Println("Listening on: ", CONN_HOST+":"+CONN_PORT)
	for {
		conn, err := li.Accept()
		for err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer
	_, er := conn.Read(buf)
	if er != nil {
		fmt.Println("Error reading: ", er)
	}

	// Send a response back to the person contacting us
	conn.Write([]byte("Message received"))
	// Close the connection when you're done with it
	conn.Close()

}
