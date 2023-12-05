package main

import (
	"encoding/binary"
	"strings"
)

type ResourceRecord struct {
	Name   string
	Type   uint16
	Class  uint16
	TTL    uint32
	Length uint16
	Data   []byte
}

func (r *ResourceRecord) DNSBinary() (data []byte) {
	labels := strings.Split(r.Name, ".")
	data = make([]byte, 0, len(r.Name)+len(labels)+10+len(r.Data))
	for _, label := range labels {
		data = append(data, byte(len(label)))
		data = append(data, []byte(label)...)
	}
	data = append(data, byte('\x00'))
	data = binary.BigEndian.AppendUint16(data, r.Type)
	data = binary.BigEndian.AppendUint16(data, r.Class)
	data = binary.BigEndian.AppendUint32(data, r.TTL)
	data = binary.BigEndian.AppendUint16(data, r.Length)
	data = append(data, r.Data...)
	return
}
