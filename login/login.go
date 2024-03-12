package login

import (
	"fmt"

	obj "irc_client/objects"
	"irc_client/pkg"

	"net"
)

type UserSession struct {
	Name        string
	token       string
	HostAddress string
}

func storeUserSession(username string, token string, hostAddress string) {
	LoginCache[username] = UserSession{Name: username, token: token, HostAddress: hostAddress}
}

var LoginCache map[string]UserSession = make(map[string]UserSession)

func ProcessLogin(connection net.Conn) {
	inputUsername, inputPassword := pkg.ReadUser() //TODO
	var response *obj.Response = processUserInput(inputUsername, inputPassword, connection)
	if response.Msg == "SUCCESS" && response.Token != "" {
		fmt.Println("User logger in")
		storeUserSession(inputUsername, response.Token, connection.LocalAddr().String())
		ProcessMessaging(connection, response.Token)
	} else if response.Msg == "FAILED" {
		fmt.Println("Login Failed, check username and password")
	} else {
		fmt.Println("Error on server, contact your admin !")
	}
}

func processUserInput(inputUsername string, inputPassword string, connection net.Conn) *obj.Response {
	user, hostAddress := pkg.CreateUserAndHostMessage(inputUsername, inputPassword)
	var userLoginData *obj.Message = &obj.Message{Command: obj.Command_Log, User: user, HostAddress: hostAddress}
	var logMsgBytes []byte = pkg.EncodeMessage(userLoginData)
	return pkg.SendRQGetRS(connection, logMsgBytes)
}
