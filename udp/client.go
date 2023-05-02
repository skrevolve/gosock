package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	udpServerAddr, err := net.ResolveUDPAddr("udp4", ":32452")
	if err != nil {
		os.Exit(1)
	}

	udpServer, err := net.DialUDP("udp", nil, udpServerAddr)
	defer udpServer.Close()
	if err != nil {
		os.Exit(2)
	}

	var inputMessage string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputMessage = scanner.Text()
		send(udpServer, inputMessage)
	}

	// os.Exit(0)
}

func send(dial *net.UDPConn, inputMessage string) {
	dial.Write([]byte(inputMessage))
}

func recv(dial *net.UDPConn) {
	var buffer[256]byte
	n, err := dial.Read(buffer[0:])
	if err != nil {
		os.Exit(3)
	}
	fmt.Println("server: ", string(buffer[0:n]))
}