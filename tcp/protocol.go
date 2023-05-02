package main

type Packet interface {
	serialize() []byte
}

type Protocol interface {
	ReadPacket(recvData []byte) (Packet, int16)
}