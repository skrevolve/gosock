package main

import "encoding/binary"

type EchoPacketHeader struct {
	TotalSize int16
	Id        int16
}

type EchoPacket struct {
	buff []byte
}

func (t *EchoPacket) Serialize() []byte {
	return t.buff
}

func (t *EchoPacket) GetLength() uint16 {
	return binary.BigEndian.Uint16(t.buff)
}

func (t *EchoPacket) GetId() uint16 {
	return binary.BigEndian.Uint16(t.buff[2:])
}

func (t *EchoPacket) GetBody() []byte {
	return t.buff[4:]
}

func NewEchoPacket(buff []byte, hasLengthField bool) *EchoPacket {
	p := &EchoPacket{}

	if hasLengthField {

		p.buff = buff

	} else {

		p.buff = make([]byte, 4 + len(buff))
		binary.BigEndian.PutUint16(p.buff[0:2], uint16(len(buff)))
		copy(p.buff[4:], buff)
	}

	return p
}

type EchoProtocol struct {}

const (
	HEADER_SIZE = 4
	MAX_PACKET_SIZE = 1024
)

func (this *EchoProtocol) ReadPacket(recvData []byte) (Packet, int16) {
	readAbleByte := int16(len(recvData))

	if readAbleByte < HEADER_SIZE {
		return nil, 0
	}

	requireDataSize := int16(binary.LittleEndian.Uint16(recvData))

	if requireDataSize > readAbleByte {
		return nil, 0
	}

	if requireDataSize > MAX_PACKET_SIZE {
		return nil, 0
	}

	ltvPacket := recvData[0:requireDataSize]
	newPacketData := make([]byte, requireDataSize)
	copy(newPacketData, ltvPacket)

	return NewEchoPacket(newPacketData, true), requireDataSize
}