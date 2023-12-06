package main

import (
	"fmt"
	"net"
)

type Connection struct {
	address net.UDPAddr
}

func NewConnection(address string) (Connection, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return Connection{}, err
	}
	return Connection{address: *udpAddr}, nil
}

func (connection *Connection) ListenToMessages(handler func(Message) Message) {
	udpConn, err := net.ListenUDP("udp", &connection.address)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()
	buf := make([]byte, 512)
	for {
		_, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
		var message = ParseDNSMessage(buf)
		var response = handler(message)
		_, err = udpConn.WriteToUDP(response.ToBytes(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
