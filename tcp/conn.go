package main

import (
	"errors"
	"net"
	"sync"
)

// error type
var (
	ErrConnClosing = errors.New("use of closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking = errors.New("read packet was blocking")
)

// Conn exposes a set of callbacks for the various events that occur a connection
type Conn struct {
	src               *Server
	tcpConn           *net.TCPConn
	extraData         interface{}
	closeOnce         sync.Once
	closeFlag         int32
	closeChan         chan struct{}
	packetSendChan    chan Packet
	packetReceiveChan chan Packet
}

//ConnCallback is an interface of methods that are used as callbacks on a connection
type ConnCallback interface {
	// OnConnect is called when the connection was accepted
	// If the return value of flase is closed
	OnConnect(*Conn) bool

	// OnMessage is called when the connection receives a packet,
	// If the return  value of false is closed
	OnMessage(*Conn, Packet) bool

	// OnClose is called when the connection closed
	OnClose(*Conn)
}

