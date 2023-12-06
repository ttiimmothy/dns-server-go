package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func answerQuestions(message Message, resolverConnection net.Conn, resolverAddress *net.UDPAddr, resolverAddrString string) []ResourceRecord {
	var answers []ResourceRecord
	for _, question := range message.Questions {
		if resolverAddrString != "" {
			resolverMessage := Message{
				Header: Header{
					ID:                  message.Header.ID,
					QR:                  0,
					OperationCode:       OPCODE_QUERY,
					AuthoritativeAnswer: 0,
					Truncation:          0,
					RecursionDesired:    0,
					RecursionAvailable:  0,
					ResponseCode:        RCODE_OK,
				},
				Questions: []Question{question},
			}
			fmt.Println("Request message:", resolverMessage)
			resolverConnection.Write(resolverMessage.ToBytes())
			deadline := time.Now().Add(15 * time.Second)
			resolverConnection.SetReadDeadline(deadline)
			buf := make([]byte, 512)
			_, err := resolverConnection.Read(buf)

			if err != nil {
				fmt.Println("Error receiving data:", err)
				break
			}

			var resolverResponse = ParseDNSMessage(buf)
			fmt.Println("Response:", resolverResponse)

			if len(resolverResponse.Answers) != 0 {
				for _, answer := range resolverResponse.Answers {
					fmt.Println("Answer:", answer.RDATA)
					answers = append(answers,
						ResourceRecord{
							Name:     question.Name,
							Type:     RR_TYPE_A,
							Class:    RR_CLASS_IN,
							TTL:      10,
							RDLENGTH: 4,
							RDATA:    answer.RDATA,
						},
					)
				}
			} else {
				answers = append(answers,
					ResourceRecord{
						Name:     question.Name,
						Type:     RR_TYPE_A,
						Class:    RR_CLASS_IN,
						TTL:      10,
						RDLENGTH: 4,
						RDATA:    "8888",
					},
				)
			}
		} else {
			answers = append(answers,
				ResourceRecord{
					Name:     question.Name,
					Type:     RR_TYPE_A,
					Class:    RR_CLASS_IN,
					TTL:      10,
					RDLENGTH: 4,
					RDATA:    "8888",
				},
			)
		}
	}
	return answers
}

func buildMessageHandler(resolverAddr net.UDPAddr, resolverConnection net.Conn, resolverAddrString string) func(Message) Message {
	return func(message Message) Message {
		var response Message
		switch message.Header.OperationCode {
		case OPCODE_QUERY:
			response = message.GenerateResponseMessage(
				RCODE_OK,
				message.Questions,
				answerQuestions(message, resolverConnection, &resolverAddr, resolverAddrString))
		default:
			response = message.GenerateResponseMessage(RCODE_NOT_IMPLEMENTED, []Question{}, []ResourceRecord{})
		}

		return response
	}
}

func main() {
	connection, err := NewConnection("127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	resolverAddrString := flag.String("resolver", "", "DNS resolver")
	flag.Parse()
	var str = string(*resolverAddrString)
	fmt.Println("Resolver addreess: ", str)
	resolverAddr, err := net.ResolveUDPAddr("udp", str)
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
	}
	resolverConnection, err := net.Dial("udp", str)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
	}
	defer resolverConnection.Close()
	connection.ListenToMessages(buildMessageHandler(*resolverAddr, resolverConnection, str))
}
