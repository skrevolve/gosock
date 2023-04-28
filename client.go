package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	port := ":8080"
	protocol := "udp4"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		os.Exit(1)
	}

	dial, err := net.DialUDP(protocol, nil, udpAddr)
	// defer dial.Close()
	if err != nil {
		os.Exit(2)
	}

	var inputMessage string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputMessage = scanner.Text()
		sendMessage(dial, inputMessage)
	}

	// os.Exit(0)
}

func sendMessage(dial *net.UDPConn, inputMessage string) {

	dial.Write([]byte(inputMessage))

	var buffer[256]byte
	n, err := dial.Read(buffer[0:])
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	fmt.Println("server: ", string(buffer[0:n]))
}