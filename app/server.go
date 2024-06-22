package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//binding to port
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
	   fmt.Println("Failed to bind to port 4221")
	   os.Exit(1)
	}
	//start to accepting connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
	 	os.Exit(1)
	}

	//read connection into byte
	req := make([]byte, 1024)
	conn.Read(req)

	path := strings.Split(conn, " ")[1]
	if path == "/" {
		connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Split(path, "/")[1] == "echo" {
		message := strings.Split(path, "/")[2]
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)))
	} else {
		connection.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
