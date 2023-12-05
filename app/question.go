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

func (q *Question) DNSBinaryByte(data *[]byte) {
	var i, dataLen int
	flag := false
	tmpName := []byte{}
	for i, dataLen = 0, len(*data); i < dataLen; {
		length := (*data)[i]
		i++
		if length == 0 {
			q.Name = string(tmpName)
			break
		}
		if flag {
			tmpName = append(tmpName, byte('.'))
		}
		flag = true
		for j := byte(0); j < length; j++ {
			tmpName = append(tmpName, (*data)[i])
			i++
		}
	}
	q.Type = binary.BigEndian.Uint16((*data)[i:])
	i += 2
	q.Class = binary.BigEndian.Uint16((*data)[i:])
	i += 2
	*data = (*data)[i:]
}
