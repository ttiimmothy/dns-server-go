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

func (h *Header) DNSBinary() (data []byte) {
	data = make([]byte, 12)
	binary.BigEndian.PutUint16(data[0:], h.ID)
	binary.BigEndian.PutUint16(data[2:],
		uint16(h.QR)<<15|
			uint16(h.OPCODE)<<11|
			uint16(h.AA)<<10|
			uint16(h.TC)<<9|
			uint16(h.RD)<<8|
			uint16(h.RA)<<7|
			uint16(h.Z)<<4|
			uint16(h.RCODE),
	)
	binary.BigEndian.PutUint16(data[4:], h.QDCOUNT)
	binary.BigEndian.PutUint16(data[6:], h.ANCOUNT)
	binary.BigEndian.PutUint16(data[8:], h.NSCOUNT)
	binary.BigEndian.PutUint16(data[10:], h.ARCOUNT)
	return
}

func (h *Header) DNSBinaryByte(data []byte) {
	h.ID = binary.BigEndian.Uint16(data[:2])
	h.QR = data[2] >> 7 & 1
	h.OPCODE = data[2] >> 3 & 7
	h.AA = data[2] >> 2 & 1
	h.TC = data[2] >> 1 & 1
	h.RD = data[2] & 1
	h.RA = data[3] >> 7 & 1
	h.Z = data[3] >> 4 & 7
	h.RCODE = data[3] & 15
	h.QDCOUNT = binary.BigEndian.Uint16(data[4:6])
	h.ANCOUNT = binary.BigEndian.Uint16(data[6:8])
	h.NSCOUNT = binary.BigEndian.Uint16(data[8:10])
	h.ARCOUNT = binary.BigEndian.Uint16(data[10:12])
}
