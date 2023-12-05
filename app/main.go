package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)
	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		receivedMessage := Message{}
		receivedMessage.DNSBinaryByte(buf)
		RCODEfromOPCODE := func(OPCODE byte) (RCODE byte) {
			if OPCODE == 0 {
				RCODE = 0
			} else {
				RCODE = 4
			}
			return
		}
		respondedMessage := Message{
			header: Header{
				ID:      receivedMessage.header.ID,
				QR:      1,
				OPCODE:  receivedMessage.header.OPCODE,
				AA:      0,
				TC:      0,
				RD:      receivedMessage.header.RD,
				RA:      0,
				Z:       0,
				RCODE:   RCODEfromOPCODE(receivedMessage.header.OPCODE),
				QDCOUNT: 1,
				ANCOUNT: 1,
				NSCOUNT: 0,
				ARCOUNT: 0,
			},
			questions: []Question{
				{
					Name:  "codecrafters.io",
					Type:  1,
					Class: 1,
				},
			},
			answer: []ResourceRecord{
				{
					Name:   "codecrafters.io",
					Type:   1,
					Class:  1,
					TTL:    60,
					Length: 4,
					Data:   []byte{8, 8, 8, 8},
				},
			},
		}

		response := respondedMessage.DNSBinary()
		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
