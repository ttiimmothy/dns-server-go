package main

type QTYPE uint16

const (
	QTYPE_A QTYPE = 1
)

type QCLASS uint16

const (
	QCLASS_IN QCLASS = 1
)

type Question struct {
	Name  []string
	Type  QTYPE
	Class QCLASS
}

func (question *Question) ToBytes() []byte {
	var bytes []byte
	for _, label := range question.Name {
		bytes = append(bytes, byte(len(label)))
		bytes = append(bytes, []byte(label)...)
	}
	bytes = append(bytes, 0)
	bytes = append(bytes, Uint16ToBytes(uint16(question.Type))...)
	bytes = append(bytes, Uint16ToBytes(uint16(question.Class))...)
	return bytes
}
