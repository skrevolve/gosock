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

	udp, err := net.ListenUDP("udp", udpAddr)
	defer udp.Close()
	if err != nil {
		os.Exit(2)
	}

	for {
		var buffer [256]byte
		n, addr, err := udp.ReadFromUDP(buffer[0:])
		if err != nil {
			continue
		}

		getMessage := fmt.Sprintf("hello %s", string(buffer[0:n]))
		udp.WriteToUDP([]byte(getMessage), addr)
	}

}