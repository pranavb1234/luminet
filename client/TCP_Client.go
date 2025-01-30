package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the local TCP server on port 8000
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Read input from the user and send to server
	fmt.Println("Connected! Type your message:")

	for {
		// Read input from the user
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		// Send message to server
		conn.Write([]byte(message))

		// Read server response
		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Server replied:", reply)
	}
}
