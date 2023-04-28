package tcp

import (
	"fmt"
	"net"
)

func Init() {

	// connect tcp
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(":: TCP Socket Connection Error ::")
	}
	defer conn.Close()

	// read
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	// write
	conn.Write([]byte("Hello from client"))
}