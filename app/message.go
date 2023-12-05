package main

type Message struct {
	header    Header
	questions []Question
	answer    []ResourceRecord
}

func (m *Message) DNSBinary() (data []byte) {
	data = m.header.DNSBinary()
	for i := range m.questions {
		data = append(data, m.questions[i].DNSBinary()...)
	}
	for i := range m.answer {
		data = append(data, m.answer[i].DNSBinary()...)
	}
	return
}

func (m *Message) DNSBinaryByte(data []byte) {
	m.header.DNSBinaryByte(data[:12])
}
