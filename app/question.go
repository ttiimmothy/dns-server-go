package main

import (
	"encoding/binary"
	"strings"
)

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func (q *Question) DNSBinary() (data []byte) {
	labels := strings.Split(q.Name, ".")
	data = make([]byte, 0, len(q.Name)+len(labels)+4)
	for _, label := range labels {
		data = append(data, byte(len(label)))
		data = append(data, []byte(label)...)
	}
	data = append(data, byte('\x00'))
	data = binary.BigEndian.AppendUint16(data, q.Type)
	data = binary.BigEndian.AppendUint16(data, q.Class)
	return
}
