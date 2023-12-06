package main

type RR_TYPE = uint16

const (
	RR_TYPE_A RR_TYPE = 1
)

type RR_CLASS uint16

const (
	RR_CLASS_IN RR_CLASS = 1
)

type ResourceRecord struct {
	Name     []string
	Type     RR_TYPE
	Class    RR_CLASS
	TTL      uint32
	RDLENGTH uint16
	RDATA    string
}

func (question *ResourceRecord) ToBytes() []byte {
	var bytes []byte
	for _, label := range question.Name {
		bytes = append(bytes, byte(len(label)))
		bytes = append(bytes, []byte(label)...)
	}
	bytes = append(bytes, 0)
	bytes = append(bytes, Uint16ToBytes(question.Type)...)
	bytes = append(bytes, Uint16ToBytes(uint16(question.Class))...)
	bytes = append(bytes, Uint32ToBytes(question.TTL)...)
	bytes = append(bytes, Uint16ToBytes(question.RDLENGTH)...)
	bytes = append(bytes, []byte(question.RDATA)...)
	return bytes
}
