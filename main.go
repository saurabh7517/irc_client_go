package main

import (
	"bufio"
	"fmt"
	"irc_client/login"
	obj "irc_client/objects"
	"irc_client/pkg"
	register "irc_client/registration"
	"log"
	"net"
	"os"
	"strconv"
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

	// localIPAddress = () TODO
	var pingMessage *obj.Message = &obj.Message{Command: obj.Command_Ping}
	var msgBytes []byte = pkg.EncodeMessage(pingMessage)

	connnection.Write(msgBytes)

	var incomingBytes []byte = make([]byte, 512)

	incomingByteLen, err := connnection.Read(incomingBytes)
	if err != nil {
		log.Println("Error reading from the connection buffer")
	}
	log.Printf("Read %d from connection buffer", incomingByteLen)
	var incomMsg *obj.Message = pkg.DecodeMessage(incomingBytes)
	log.Println("Server :: ", incomMsg.Command)

	if incomMsg.Command == obj.Command_Pong {
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
		var regMsgBytes []byte = register.ProcessRegistration(reader)
		var response *obj.Response = pkg.SendRQGetRS(connnection, regMsgBytes)
		if response.Msg == "CREATED" {
			fmt.Println("New User Created")
		} else if response.Msg == "ALREADY_EXISTS" {
			fmt.Println("User already exists")
		} else {
			fmt.Println("Error creating user on server, Try again !!")
		}
	case 2:
		var logMsgBytes []byte = login.ProcessLogin(reader)
		var response *obj.Response = pkg.SendRQGetRS(connnection, logMsgBytes)
		if response.Msg == "SUCCESS" && response.Token != "" {
			fmt.Println("User logger in")
			pkg.ProcessMessaging(connnection, response.Token)
		} else if response.Msg == "FAILED" {
			fmt.Println("Login Failed, check username and password")
		} else {
			fmt.Println("Error on server, contact your admin !")
		}

	case 3:
		os.Exit(1)
	}

}
