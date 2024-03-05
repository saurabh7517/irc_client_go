package pkg

import (
	obj "irc_client/objects"
	"irc_client/pkg"
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
	var tokenMsg *obj.TokenMsg = pkg.CreateTokenMessage(token)
	pkg.EncodeMessage(tokenMsg)

}
