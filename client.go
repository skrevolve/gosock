package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":1800"

	udpAddr, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		os.Exit(1)
	}

	dial, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		os.Exit(2)
	}

	dial.Write([]byte("im calvin"))

	var buffer[256]byte
	n, err := dial.Read(buffer[0:])
	if err != nil {
		os.Exit(3)
	}

	fmt.Println("Response from server: ", string(buffer[0:n]))

	dial.Close()
	os.Exit(0)
}