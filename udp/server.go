package udp

import (
	"fmt"
	"net"
	"os"
)

func main() {

	port := ":8080"
	protocol := "udp4"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		fmt.Println("Wrong address")
		os.Exit(1)
	}

	udpConn, err := net.ListenUDP(protocol, udpAddr)
	//defer udpConn.Close()
	if err != nil {
		os.Exit(2)
	}

	for {
		getClientMessage(udpConn)
	}
}

func getClientMessage(conn *net.UDPConn) {
	var buffer [256]byte
	n, addr, _ := conn.ReadFromUDP(buffer[0:])
	// if err != nil { }
	clientMessage := string(buffer[0:n])
	fmt.Println("client : " + clientMessage)

	serverMessage := "ok"
	conn.WriteToUDP([]byte(serverMessage), addr)
}