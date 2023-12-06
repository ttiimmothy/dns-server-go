package main

import (
	"encoding/binary"
)

type Message struct {
	Header    Header
	Questions []Question
	Answers   []ResourceRecord
}

func parseLabels(bytes []byte, startIndex int) ([]string, int) {
	var labels []string
	var currentByteIndex = startIndex
	var currentByte = bytes[currentByteIndex]
	for currentByte > 0 {
		if currentByte&0xC0>>6 == 3 {
			var referencedByteIndex = int(binary.BigEndian.Uint16([]byte{currentByte & 0x3F, bytes[currentByteIndex+1]}))
			var referencedLabels, _ = parseLabels(bytes, referencedByteIndex)
			labels = append(labels, referencedLabels...)
			currentByteIndex += 2
			break
		}
		currentByteIndex += 1
		var labelBytes = bytes[currentByteIndex : currentByteIndex+int(currentByte)]
		labels = append(labels, string(labelBytes))
		currentByteIndex += int(currentByte)
		currentByte = bytes[currentByteIndex]
	}
	return labels, currentByteIndex
}

func ParseDNSMessage(bytes []byte) Message {
	var currentByteIndex = 12
	var rawHeader = RawDNSHeaderFromBytes(bytes[:currentByteIndex])
	var questions []Question
	for i := 0; i < int(rawHeader.QDCOUNT); i++ {
		var labels, labelsEndIndex = parseLabels(bytes, currentByteIndex)
		currentByteIndex = labelsEndIndex + 1
		questions = append(
			questions,
			Question{
				Name:  labels,
				Type:  QTYPE(binary.BigEndian.Uint16(bytes[currentByteIndex : currentByteIndex+2])),
				Class: QCLASS(binary.BigEndian.Uint16(bytes[currentByteIndex+2 : currentByteIndex+4])),
			},
		)
		currentByteIndex += 4
	}

	var answers []ResourceRecord
	for i := 0; i < int(rawHeader.ANCOUNT); i++ {
		var labels, labelsEndIndex = parseLabels(bytes, currentByteIndex)
		currentByteIndex = labelsEndIndex + 1
		rr_type := RR_TYPE(binary.BigEndian.Uint16(bytes[currentByteIndex : currentByteIndex+2]))
		rr_class := RR_CLASS(binary.BigEndian.Uint16(bytes[currentByteIndex+2 : currentByteIndex+4]))
		rr_ttl := binary.BigEndian.Uint32(bytes[currentByteIndex+4 : currentByteIndex+8])
		rdlen := binary.BigEndian.Uint16(bytes[currentByteIndex+8 : currentByteIndex+10])
		rdata := string(bytes[currentByteIndex+10 : currentByteIndex+10+int(rdlen)])
		answers = append(
			answers,
			ResourceRecord{
				Name:     labels,
				Type:     rr_type,
				Class:    rr_class,
				TTL:      rr_ttl,
				RDLENGTH: rdlen,
				RDATA:    rdata,
			},
		)
		currentByteIndex += 4
	}
	return Message{
		Header: Header{
			ID:                  rawHeader.ID,
			QR:                  rawHeader.QR,
			OperationCode:       OPCODE(rawHeader.OPCODE),
			AuthoritativeAnswer: rawHeader.AA,
			Truncation:          rawHeader.TC,
			RecursionDesired:    rawHeader.RD,
			RecursionAvailable:  rawHeader.RA,
			ResponseCode:        RCODE(rawHeader.RCODE),
		},
		Questions: questions,
		Answers:   answers,
	}
}

func (message *Message) ToBytes() []byte {
	var bytes []byte
	var rawHeader = RawDNSHeader{
		ID:      message.Header.ID,
		QR:      message.Header.QR,
		OPCODE:  byte(message.Header.OperationCode),
		AA:      message.Header.AuthoritativeAnswer,
		TC:      message.Header.Truncation,
		RD:      message.Header.RecursionDesired,
		RA:      message.Header.RecursionAvailable,
		Z:       0,
		RCODE:   byte(message.Header.ResponseCode),
		QDCOUNT: uint16(len(message.Questions)),
		ANCOUNT: uint16(len(message.Answers)),
		NSCOUNT: 0,
		ARCOUNT: 0,
	}
	bytes = append(bytes, rawHeader.ToBytes()...)
	for _, question := range message.Questions {
		bytes = append(bytes, question.ToBytes()...)
	}
	for _, answer := range message.Answers {
		bytes = append(bytes, answer.ToBytes()...)
	}
	return bytes
}

func (message *Message) GenerateResponseMessage(response_code RCODE, questions []Question, answers []ResourceRecord) Message {
	var responseHeader = Header{
		ID:                  message.Header.ID,
		QR:                  1,
		OperationCode:       message.Header.OperationCode,
		AuthoritativeAnswer: 0,
		Truncation:          0,
		RecursionDesired:    message.Header.RecursionDesired,
		RecursionAvailable:  0,
		ResponseCode:        response_code,
	}
	return Message{
		Header:    responseHeader,
		Questions: questions,
		Answers:   answers,
	}
}
