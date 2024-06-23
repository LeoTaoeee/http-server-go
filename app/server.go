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

	if path == "/" {
		//default 200OK
		connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Split(path, "/")[1] == "echo" {
		//echo request
		message := strings.Split(path, "/")[2]
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)))
	}else if strings.Split(path, "/")[1] == "user-agent"{
		//user-agent
		temp := strings.Split(request, ":")[2]
		message := strings.Split(temp,"\r\n")[0]
		connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)))
	}else {
		//invalid 404
		connection.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}