package main
import (
	"fmt"
	"net"
	"os"
	"strings"
)
func main() {
	//bind to port 4221
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close() //close when everything finished

	//accept connection
	for {
		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(connection)
	}
}
func handleConnection(connection net.Conn) {
	defer connection.Close() //close when everything finished

	//read conn into buffer
	requestBuffer := make([]byte, 1024)
	n, err := connection.Read(requestBuffer)
	if err != nil {
		fmt.Println("Failed to read the request:", err)
		return
	}

	//console testing
	fmt.Printf("Request: %s\n", requestBuffer[:n])

	request := string(requestBuffer[:n])

	//retrieve url path
	path := strings.Split(request, " ")[1]

	//default 200OK
	if path == "/" {
		connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	}
	//echo request
	else if strings.Split(path, "/")[1] == "echo" {
		message := strings.Split(path, "/")[2]
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)))
	}
	//invalid 404
	else {
		connection.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}