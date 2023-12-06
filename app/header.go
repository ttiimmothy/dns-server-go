package main

import (
	"encoding/binary"
)

type OPCODE byte

const (
	OPCODE_QUERY  = 0
	OPCODE_IQUERY = 1
	OPCODE_STATUS = 2
)

type RCODE byte

const (
	RCODE_OK              = 0
	RCODE_NOT_IMPLEMENTED = 4
)

type Header struct {
	ID                  uint16
	QR                  byte
	OperationCode       OPCODE
	AuthoritativeAnswer byte
	Truncation          byte
	RecursionDesired    byte
	RecursionAvailable  byte
	ResponseCode        RCODE
}
type RawDNSHeader struct {
	ID      uint16
	QR      byte
	OPCODE  byte
	AA      byte
	TC      byte
	RD      byte
	RA      byte
	Z       byte
	RCODE   byte
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

func RawDNSHeaderFromBytes(bytes []byte) RawDNSHeader {
	return RawDNSHeader{
		ID:      binary.BigEndian.Uint16(bytes[:2]),
		QR:      bytes[2] & 0x80 >> 7,
		OPCODE:  bytes[2] & 0x78 >> 3,
		AA:      bytes[2] & 0x4 >> 2,
		TC:      bytes[2] & 0x2 >> 1,
		RD:      bytes[2] & 0x1,
		RA:      bytes[3] & 0x80 >> 7,
		Z:       bytes[3] & 0x70 >> 4,
		RCODE:   bytes[3] & 0xF,
		QDCOUNT: binary.BigEndian.Uint16(bytes[4:6]),
		ANCOUNT: binary.BigEndian.Uint16(bytes[6:8]),
		NSCOUNT: binary.BigEndian.Uint16(bytes[8:10]),
		ARCOUNT: binary.BigEndian.Uint16(bytes[10:12]),
	}
}

func (header *RawDNSHeader) ToBytes() []byte {
	var bytes []byte
	bytes = append(bytes, Uint16ToBytes(header.ID)...)
	bytes = append(bytes,
		byte(header.QR<<7|header.OPCODE<<3|header.AA<<2|header.TC<<1|header.RD),
		byte(header.RA<<7|header.Z<<4|header.RCODE))
	bytes = append(bytes, Uint16ToBytes(header.QDCOUNT)...)
	bytes = append(bytes, Uint16ToBytes(header.ANCOUNT)...)
	bytes = append(bytes, Uint16ToBytes(header.NSCOUNT)...)
	bytes = append(bytes, Uint16ToBytes(header.ARCOUNT)...)
	return bytes
}
