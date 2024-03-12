package register

import (
	"fmt"
	obj "irc_client/objects"
	"irc_client/pkg"
	"net"
)

func ProcessRegistration(connection net.Conn) {
	initiateRegistrationProcess(connection)
}

func initiateRegistrationProcess(connection net.Conn) {
	inputUser, inputPassword := pkg.ReadUser()
	user, hostAddress := pkg.CreateUserAndHostMessage(inputUser, inputPassword)
	var userRegData *obj.Message = &obj.Message{Command: obj.Command_Reg, User: user, HostAddress: hostAddress, PrivateMsg: nil}
	var regMsgBytes []byte = pkg.EncodeMessage(userRegData)
	var response *obj.Response = pkg.SendRQGetRS(connection, regMsgBytes)
	if response.Msg == "CREATED" {
		fmt.Println("New User Created")
	} else if response.Msg == "ALREADY_EXISTS" {
		fmt.Println("User already exists")
	} else {
		fmt.Println("Error creating user on server, Try again !!")
	}
}
