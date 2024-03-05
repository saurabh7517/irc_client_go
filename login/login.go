package login

import (
	"bufio"
	obj "irc_client/objects"
	"irc_client/pkg"
)

func ProcessLogin(reader *bufio.Reader) []byte {
	inputUser, inputPassword := pkg.ReadUser(reader) //TODO
	user, hostAddress := pkg.CreateUserAndHostMessage(inputUser, inputPassword)
	var userLoginData *obj.Message = &obj.Message{Command: obj.Command_Log, User: user, HostAddress: hostAddress}
	return pkg.EncodeMessage(userLoginData)
}
