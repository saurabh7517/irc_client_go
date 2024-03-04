package main

import (
	"bufio"
	"fmt"
	obj "irc_client/objects"
	comm "irc_client/pkg"
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
	var pingMessage *obj.Message = &obj.Message{Command: obj.Command_Ping}
	var msgBytes []byte = encodeMessage(pingMessage)

	connnection.Write(msgBytes)

	var incomingBytes []byte = make([]byte, 512)

	incomingByteLen, err := connnection.Read(incomingBytes)
	if err != nil {
		log.Println("Error reading from the connection buffer")
	}
	log.Printf("Read %d from connection buffer", incomingByteLen)
	var incomMsg *obj.Message = decodeMessage(incomingBytes)
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
		var regMsgBytes []byte = processRegistration(reader)
		var response *obj.Response = comm.SendRQGetRS(connnection, regMsgBytes)
		if response.Msg == "CREATED" {
			fmt.Println("New User Created")
		} else if response.Msg == "ALREADY_EXISTS" {
			fmt.Println("User already exists")
		} else {
			fmt.Println("Error creating user on server, Try again !!")
		}
	case 2:
		var logMsgBytes []byte = loginUser(reader)
		var response *obj.Response = comm.SendRQGetRS(connnection, logMsgBytes)
		if response.Msg == "SUCCESS" {
			fmt.Println("User logger in")
			processUser()
		} else if response.Msg == "FAILED" {
			fmt.Println("Login Failed, check username and password")
		} else {
			fmt.Println("Error on server, contact your admin !")
		}

	case 3:
		os.Exit(1)
	}

}

func processRegistration(reader *bufio.Reader) []byte {
	inputUser, inputPassword := readUser(reader)
	user, hostAddress := createUserAndHostMessage(inputUser, inputPassword)
	var userRegData *obj.Message = &obj.Message{Command: obj.Command_Reg, User: user, HostAddress: hostAddress}
	return encodeMessage(userRegData)
}

func loginUser(reader *bufio.Reader) []byte {
	inputUser, inputPassword := readUser(reader)
	user, hostAddress := createUserAndHostMessage(inputUser, inputPassword)
	var userLoginData *obj.Message = &obj.Message{Command: obj.Command_Log, User: user, HostAddress: hostAddress}
	return encodeMessage(userLoginData)
}

func createUserAndHostMessage(username string, password string) (*obj.User, *obj.HostAddress) {
	var user *obj.User = &obj.User{Username: username, Password: password}
	var localAddress string
	if localIPAddress != "" {
		localAddress = localIPAddress
	} else {
		localAddress = getLocalIPAddress()
	}
	var hostAddress *obj.HostAddress = &obj.HostAddress{HostIp: localAddress, HostPort: ""}

	return user, hostAddress
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

func encodeMessage(message *obj.Message) []byte {
	msgBytes, err := proto.Marshal(message)
	if err != nil {
		log.Println("Error in marshaling data")
	}
	return msgBytes
}

func decodeMessage(message []byte) *obj.Message {
	var incomMsg *obj.Message = &obj.Message{}
	proto.Unmarshal(message, incomMsg)
	return incomMsg
}

func getLocalIPAddress() string {
	var ipAddr string
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Panic("Error getting network interfaces")
	}

	for _, i := range ifaces {
		if i.Flags&net.FlagRunning != 0 && !strings.Contains(i.Name, "Loopback") && i.MTU != -1 {
			addrs, err := i.Addrs()
			// handle err
			if err != nil {
				log.Print("Error getting network address of interface")
			}
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					ip := v.IP.To4()
					if ip != nil {
						ipAddr = ip.String()
					}
				}
			}
		}
	}
	return ipAddr
}