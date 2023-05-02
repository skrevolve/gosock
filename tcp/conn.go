package main

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// error type
var (
	ErrConnClosing = errors.New("use of closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking = errors.New("read packet was blocking")
)

// Conn exposes a set of callbacks for the various events that occur a connection
type Conn struct {
	srv               *Server
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

// newConn returns a wrapper of raw tcpConn
func newConn(conn *net.TCPConn, srv *Server) *Conn {
	return &Conn{
		srv:               srv,
		tcpConn:           conn,
		closeChan:         make(chan struct{}),
		packetSendChan:    make(chan Packet, srv.config.PacketSendChanLimit),
		packetReceiveChan: make(chan Packet, srv.config.PacketReceiveChanLImit),
	}
}

// GetExtraData gets the extra data from the Conn
func (c *Conn) GetExtraData() interface{} {
	return c.extraData
}

// PutExtraData puts the extra data with th Conn
func (c *Conn) PutExtraData(data interface{}) {
	c.extraData = data
}

// GetRawConn returns the raw net.TCPConn from Conn
func (c *Conn) GetRawConn() *net.TCPConn {
	return c.tcpConn
}

// Close closes the connection
func (c *Conn) Close() {
	c.closeOnce.Do(func() {
		atomic.StoreInt32(&c.closeFlag, 1)
		close(c.closeChan)
		close(c.packetSendChan)
		close(c.packetReceiveChan)
		c.tcpConn.Close()
		c.srv.callback.OnClose(c)
	})
}

// IsClosed indicates whether or not the connection is closed
func (c *Conn) IsClosed() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

// AsyncWritePacket async writes a packet, this method will never block
func (c *Conn) AsyncWritePacket(p Packet, timeout time.Duration) (err error) {
	if c.IsClosed() {
		return ErrConnClosing
	}

	defer func() {
		if e := recover(); e != nil {
			err = ErrConnClosing
		}
	}()

	if timeout == 0 {

		select {
		case c.packetSendChan <- p:
			return nil

		default:
			return ErrWriteBlocking
		}

	} else {

		select {
		case c.packetSendChan <- p:
			return nil

		case <- c.closeChan:
			return ErrConnClosing

		case <-time.After(timeout):
			return ErrWriteBlocking
		}
	}
}

const MAX_RECEIVE_BUFFER_SIZE = 4096

// Do it
func (c *Conn) Do() {
	if !c.srv.callback.OnConnect(c) {
		return
	}

	asyncDo(c.handleLoop, c.srv.waitGroup)
	asyncDo(c.readLoop, c.srv.waitGroup)
	asyncDo(c.writeLoop, c.srv.waitGroup)
}

func (c *Conn) handleLoop() {
	defer func() {
		recover()
		c.Close()
	}()
}

func (c *Conn) readLoop() {
	defer func() {
		recover()
		c.Close()
	}()

	var startRecvPos int16 = 0
	receiveBuff := make([]byte, MAX_RECEIVE_BUFFER_SIZE)

	for {
		select {
		case <-c.srv.exitChan:
			return

		case <-c.closeChan:
			return

		default:
		}

		recvLen, err := c.tcpConn.Read(receiveBuff[startRecvPos:])
		if err != nil {
			return
		}

		readAbleLen := startRecvPos + int16(recvLen)
		var readPos int16 = 0

		for {
			readableBuffer := receiveBuff[readPos:readAbleLen]
			p, pSize := c.srv.protocol.ReadPacket(readableBuffer)
			if p != nil {

				c.packetReceiveChan <- p
				readPos += pSize

			} else {

				remainSize := readAbleLen - readPos
				if remainSize > 0 {
					copy(receiveBuff, receiveBuff[readPos:readAbleLen])
				}
				startRecvPos = remainSize
				break;
			}
		}

	}
}

func (c *Conn) writeLoop() {
	defer func() {
		recover()
		c.Close()
	}()

	for {
		select {
		case <-c.srv.exitChan:
			return

		case <-c.closeChan:
			return

		case p := <-c.packetSendChan:
			if c.IsClosed() {
				return
			}
			if _, err := c.tcpConn.Write(p.Serialize()); err != nil {
				return
			}
		}
	}
}

func asyncDo(fn func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		fn()
		wg.Done()
	}()
}