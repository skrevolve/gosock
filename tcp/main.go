package main

import (
	"log"
	"net"
)

type Callback struct{}

func (this *Callback) OnConnect(c *Conn) bool {
	//addr :=
}

func main() {

	// create tcp listener
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":32452")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// creates a server
	config := &Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLImit: 20,
	}
	srv := NewServer(config, &Callback{}, &EchoProtocol{})
	conn, err := listener.AcceptTCP() // wait
}

func checkError(err error) {
	if err != nil { log.Fatal(err) }
}