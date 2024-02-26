package main

import (
	"fmt"
	"irc_client/pkg"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

const SERVER_HOST = "127.0.0.1"
const SERVER_PORT = "8080"
const SERVER_TYPE = "tcp4"

func main() {
	var address string = fmt.Sprintf("%v:%v", SERVER_HOST, SERVER_PORT)
	connnection, err := net.Dial(SERVER_TYPE, address)
	if err != nil {
		log.Fatalf("Cannot establish connection with Server :: %v", address)
	}
	defer connnection.Close()
	// initialize the size of buffer 512 bytes
	// send
	// var msgBytes []byte = []byte("Hello from client")
	var pingMessage *pkg.Message = &pkg.Message{Command: pkg.Command_Ping}
	var msgBytes []byte = encodeMessage(pingMessage)

	connnection.Write(msgBytes)

	var incomingBytes []byte = make([]byte, 512)

	incomingByteLen, err := connnection.Read(incomingBytes)
	if err != nil {
		log.Println("Error reading from the connection buffer")
	}
 	log.Printf("Read %d from connection buffer", incomingByteLen)
	var incomMsg *pkg.Message = decodeMessage(incomingBytes)
	log.Println("Server :: ", incomMsg.Command)

}

func encodeMessage(message *pkg.Message) []byte {
	msgBytes, err := proto.Marshal(message)
	if err != nil {
		log.Println("Error in marshaling data")
	}
	return msgBytes
}

func decodeMessage(message []byte) *pkg.Message {
	var incomMsg *pkg.Message = &pkg.Message{}
	proto.Unmarshal(message, incomMsg)
	return incomMsg
}
