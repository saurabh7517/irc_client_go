package register

import (
	"bufio"
	obj "irc_client/objects"
	"irc_client/pkg"
)

func ProcessRegistration(reader *bufio.Reader) []byte {
	inputUser, inputPassword := pkg.ReadUser(reader)
	user, hostAddress := pkg.CreateUserAndHostMessage(inputUser, inputPassword)
	var userRegData *obj.Message = &obj.Message{Command: obj.Command_Reg, User: user, HostAddress: hostAddress, Token: "", PrivateMsg: nil}
	return pkg.EncodeMessage(userRegData)
}
