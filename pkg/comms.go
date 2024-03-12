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
	proto.Unmarshal(rSBytes, response)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Println("Error un-marshalling the request bytes from server")
	// }
	return response
}

func SendRQGetActiveUsers(connection net.Conn, rQMsgBytes []byte) *obj.ActiveUser {
	var rSBytes []byte = serRQDeRS(connection, rQMsgBytes)
	var activeUsers *obj.ActiveUser = &obj.ActiveUser{}
	err := proto.Unmarshal(rSBytes, activeUsers)
	if err != nil {
		log.Println("Error un-marshalling the request bytes from server")
	}
	return activeUsers
}

func serRQDeRS(connection net.Conn, message []byte) []byte {
	var dataBytes []byte
	// send message
	if sendToServer(connection, message) {
		// receive response
		bytesFromServer, err := receiveFromServer(connection)
		if err != nil {
			log.Print("Error from server sending response")
		}
		dataBytes = bytesFromServer
	}
	return dataBytes
}

func sendToServer(connection net.Conn, message []byte) bool {
	len, err := connection.Write(message)
	if err != nil {
		var errorMsg = fmt.Sprint("Error writing to connection buffer")
		log.Println(errorMsg)
		return false
	}
	log.Printf("Written %d bytes to the connection buffer", len)
	return true
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
