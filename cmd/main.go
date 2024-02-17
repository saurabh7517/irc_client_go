package main

import (
	"fmt"
	"log"
	"net"
)

const SERVER_HOST = "192.168.3.6"
const SERVER_PORT = "9988"
const SERVER_TYPE = "tcp"

func main() {
	var address string = fmt.Sprintf("%v:%v", SERVER_HOST, SERVER_PORT)
	connnection, err := net.Dial(SERVER_TYPE, address)
	if err != nil {
		log.Fatalf("Cannot establish connection with Server :: %v", address)
	}
	defer connnection.Close()
	// initialize the size of buffer 512 bytes
	// send
	var msgBytes []byte = []byte("Hello from client")
	connnection.Write(msgBytes)
}
