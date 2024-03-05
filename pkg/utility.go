package pkg

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	obj "irc_client/objects"

	"google.golang.org/protobuf/proto"
)

func ReadUser(reader *bufio.Reader) (string, string) {

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

// TODO change this logic
func CreateUserAndHostMessage(username string, password string) (*obj.User, *obj.HostAddress) {
	var user *obj.User = &obj.User{Username: username, Password: password}
	var localAddress string = getLocalIPAddress()
	// if localIPAddress != "" {
	// 	localAddress = localIPAddress
	// } else {
	// 	localAddress = getLocalIPAddress()
	// }
	var hostAddress *obj.HostAddress = &obj.HostAddress{HostIp: localAddress, HostPort: ""}

	return user, hostAddress
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

func EncodeMessage(message *obj.Message) []byte {
	msgBytes, err := proto.Marshal(message)
	if err != nil {
		log.Println("Error in marshaling data")
	}
	return msgBytes
}

func DecodeMessage(message []byte) *obj.Message {
	var incomMsg *obj.Message = &obj.Message{}
	proto.Unmarshal(message, incomMsg)
	return incomMsg
}
