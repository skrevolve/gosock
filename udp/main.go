package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp4", ":32452")
	if err != nil {
		os.Exit(1)
	}

	server, err := net.ListenUDP("udp", udpAddr)
	defer server.Close()
	if err != nil {
		os.Exit(2)
	}

	for {
		recvAndBroadcast(server)
	}
}

func recvAndBroadcast(server *net.UDPConn) {
	var buffer [256]byte
	n, addr, _ := server.ReadFromUDP(buffer[0:])
	// if err != nil { }
	clientMessage := string(buffer[0:n])
	fmt.Print(addr)
	fmt.Println(" : " + clientMessage)

	serverMessage := "ok"
	server.WriteToUDP([]byte(serverMessage), addr)
}

// func send(conn *net.UDPConn) {

// }