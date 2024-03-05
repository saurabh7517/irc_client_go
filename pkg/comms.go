package pkg

import (
	"errors"
	"fmt"
	obj "irc_client/objects"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

func SendRQGetRS(connection net.Conn, rQMsgBytes []byte) *obj.Response {
	var rSBytes []byte = serRQDeRS(connection, rQMsgBytes)
	var response *obj.Response = &obj.Response{}
	err := proto.Unmarshal(rSBytes, response)
	if err != nil {
		log.Println("Error un-marshalling the request bytes from server")
	}
	return response
}

func serRQDeRS(connection net.Conn, message []byte) []byte {
	var dataBytes []byte
	// send message
	if sendToServer(connection, message) != nil {
		// receive response
		bytesFromServer, err := receiveFromServer(connection)
		if err != nil {
			log.Print("Error from server sending response")
		}
		dataBytes = bytesFromServer
	}
	return dataBytes
}

func sendToServer(connection net.Conn, message []byte) error {
	len, err := connection.Write(message)
	if err != nil {
		var errorMsg = fmt.Sprint("Error writing to connection buffer")
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}
	log.Printf("Written %d bytes to the connection buffer", len)
	return nil
}

func receiveFromServer(connnection net.Conn) ([]byte, error) {
	var incomingBytes []byte = make([]byte, 512)
	incomingByteLen, err := connnection.Read(incomingBytes)
	if err != nil {
		var errorMsg string = fmt.Sprint("Error reading from connection buffer")
		log.Println(errorMsg)
		return incomingBytes, errors.New(errorMsg)
	}
	log.Printf("Read %d from connection buffer", incomingByteLen)
	return incomingBytes, nil
}
