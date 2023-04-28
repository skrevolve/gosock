package udp

import (
	"fmt"
	"net"
)

func Init() {

	// connect udp
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(":: UDP Socket Connection Error ::")
	}
	defer conn.Close()

	// read
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	// write
	conn.Write([]byte("Hello from client"))
}