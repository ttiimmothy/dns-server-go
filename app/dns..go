package main

import "encoding/binary"

type Header struct {
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

func (h *Header) Tobytes() []byte {
	header := make([]byte, 12)
	binary.BigEndian.PutUint16(header[0:], h.ID)
	binary.BigEndian.PutUint16(header[2:],
		uint16(h.QR)<<15|
			uint16(h.OPCODE)<<11|
			uint16(h.AA)<<10|
			uint16(h.TC)<<9|
			uint16(h.RD)<<8|
			uint16(h.RA)<<7|
			uint16(h.Z)<<4|
			uint16(h.RCODE),
	)
	binary.BigEndian.PutUint16(header[4:], h.QDCOUNT)
	binary.BigEndian.PutUint16(header[6:], h.ANCOUNT)
	binary.BigEndian.PutUint16(header[8:], h.NSCOUNT)
	binary.BigEndian.PutUint16(header[10:], h.ARCOUNT)
	return header
}
