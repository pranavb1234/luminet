package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen for connections on port 8000
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Toy Telephone (Server) is ready on port 8000...")

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Someone connected!")

		// Handle the connection in a separate function
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Read message from client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed.")
			return
		}
		fmt.Printf("Received: %s", message)

		// Send a response back
		response := "Hello from the Toy Telephone!\n"
		conn.Write([]byte(response))
	}
}
