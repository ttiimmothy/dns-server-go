package main

type Message struct {
	header    Header
	questions []Question
}

func (m *Message) MarshalBinary() (data []byte) {
	data = m.header.MarshalBinary()
	for i := range m.questions {
		data = append(data, m.questions[i].MarshalBinary()...)
	}
	return
}
