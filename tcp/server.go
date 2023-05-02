package main

import (
	"net"
	"sync"
	"time"
)

type Config struct {
	PacketSendChanLimit    uint32
	PacketReceiveChanLImit uint32
}

type Server struct {
	config    *Config         // server config
	callback  ConnCallback    // message callbacks in connection
	protocol  Protocol        // customize packet protocol
	exitChan  chan struct{}   // notify all goroutines to shutdown
	waitGroup *sync.WaitGroup // wait for all goroutines
}

// NewServer creates a server
func NewServer(config *Config, callback ConnCallback, protocol Protocol) *Server {
	return &Server{
		config: config,
		callback: callback,
		protocol: protocol,
		exitChan: make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}
}

// Start Starts service
func (s *Server) Start(listener *net.TCPListener, acceptTimeout time.Duration) {
	s.waitGroup.Add(1)
	defer func() {
		listener.Close()
		s.waitGroup.Done()
	}

	for {
		select {
		case <-s.exitChan:
			return
		default:
		}

		listener.SetDeadline(time.Now().Add(acceptTimeout))

		conn, err := listener.AcceptTCP()
	}
}