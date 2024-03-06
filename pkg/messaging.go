package pkg

import (
	obj "irc_client/objects"
	"net"
)

func ProcessMessaging(conn net.Conn, token string) {
	//list all active users
	//create a go routine which will listen for new added active users
	getAllUsers(conn, token)

}

func showActiveUsers() {

}

func showActiveChannels() {

}

func checkForJoinedUsers() {

}

func getAllUsers(conn net.Conn, token string) []string {
	var tokenMsg *obj.Message = CreateTokenMessage(token)
	var tokenMsgBytes []byte = EncodeMessage(tokenMsg)
	var activeUsers *obj.ActiveUser = SendRQGetActiveUsers(conn, tokenMsgBytes)
	return activeUsers.GetUsername()
}
