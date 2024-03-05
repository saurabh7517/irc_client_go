package pkg

import (
	obj "irc_client/objects"
	"net"
)

func ProcessMessaging(conn net.Conn, token string) {
	//list all active users
	//create a go routine which will listen for new added active users

}

func showActiveUsers() {

}

func showActiveChannels() {

}

func checkForJoinedUsers() {

}

func getAllUsers(conn net.Conn, token string) {
	var tokenMsg *obj.Message = CreateTokenMessage(token)
	EncodeMessage(tokenMsg)

}
