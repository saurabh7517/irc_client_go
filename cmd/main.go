package main

import (
	"bufio"
	"fmt"
	"irc_client/pkg"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

const SERVER_HOST = "127.0.0.1"
const SERVER_PORT = "8080"
const SERVER_TYPE = "tcp4"

var localIPAddress string

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

	localIPAddress = getLocalIPAddress()
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

	if incomMsg.Command == pkg.Command_Pong {
		log.Println("Server is up and running")
	}

	fmt.Println("You are now connected to the chat server")
	fmt.Println("1. If you are a new user, Register")
	fmt.Println("2. If you are returning user, login")
	fmt.Println("3. Exit")
	//Creating a handle for reader here
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your option")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Error reading input")
	}
	option, err := strconv.Atoi(input)
	if err != nil {
		log.Println("Enter text in not numeric")
	}
	switch option {
	case 1:
		var regMsgBytes []byte = processRegistration(reader)
		connnection.Write(regMsgBytes)
		break
	case 2:
		loginUser()
		break
	case 3:
		os.Exit(1)
	}

}

func processRegistration(reader *bufio.Reader) []byte {
	username, password := readUser(reader)

	var user *pkg.User = &pkg.User{Username: username, Password: password}
	var localAddress string
	if localIPAddress != "" {
		localAddress = localIPAddress
	} else {
		localAddress = getLocalIPAddress()
	}
	var hostAddress *pkg.HostAddress = &pkg.HostAddress{HostIp: localAddress, HostPort: ""}

	var userRegData *pkg.Message = &pkg.Message{Command: pkg.Command_Reg, User: user, HostAddress: hostAddress}

	return encodeMessage(userRegData)

}

func readUser(reader *bufio.Reader) (string, string) {

	fmt.Println("Enter username")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Error reading input")
	}
	fmt.Println("Enter password")
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Error reading input")
	}
	username = strings.Trim(username, " ")
	password = strings.Trim(password, " ")

	return username, password
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
