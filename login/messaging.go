package login

import (
	"fmt"
	obj "irc_client/objects"
	"irc_client/pkg"
	"log"
	"net"
)

func ProcessMessaging(conn net.Conn, token string) {
	//list all active users
	//create a go routine which will listen for new added active users
	var activeUsers []string = getAllUsers(conn, token)
	if len(activeUsers) > 0 {
		showActiveUsers(activeUsers)
		fmt.Println("Enter the name of the user, you want to chat with...")
		var userInput string
		fmt.Scanf("%s", userInput)
		// reader := bufio.NewReader(os.Stdin)
		// userInput, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Println("Error reading from console")
		// }
		var found bool = false
		var targetUser string = ""
		if len(userInput) > 0 {
			for _, activeUser := range activeUsers {
				if activeUser == userInput {
					targetUser = activeUser
					found = true
					break
				}
			}
		}

		if found {
			// send a message to the user on connection

			var userSession UserSession = LoginCache[token]
			createUserSession(conn, userSession.Name, targetUser, token, userSession.HostAddress)

		} else {
			log.Println("User entered is not among active user list")
		}
	}
}

func createUserSession(conn net.Conn, sourceUser string, targetUser string, token string, hostAddress string) {
	fmt.Println("Enter 'exit' to exit chat")
	var message string = ""
	for message != "exit" {
		fmt.Print("Your message : > ") //TODO own username
		// message, _ := reader.ReadString('\n')
		var message string
		fmt.Scanf("%s", message)
		var privMsgData *obj.Message = pkg.CreateMessage(token, obj.Command_PrivMsg, message, sourceUser, targetUser, hostAddress)
		var privMsgBytes []byte = pkg.EncodeMessage(privMsgData)
		var response *obj.Response = pkg.SendRQGetRS(conn, privMsgBytes)
		if response.Status == "SUCCESS" && response.Msg == "MsgSent" {
			fmt.Println("Msg sent to server")
		}
	}
}

func showActiveUsers(userList []string) {
	for _, user := range userList {
		fmt.Println(user)
	}
}

func showActiveChannels() {

}

func checkForJoinedUsers() {

}

func getAllUsers(conn net.Conn, token string) []string {
	var tokenMsg *obj.Message = pkg.CreateTokenMessage(token)
	var tokenMsgBytes []byte = pkg.EncodeMessage(tokenMsg)
	var activeUsers *obj.ActiveUser = pkg.SendRQGetActiveUsers(conn, tokenMsgBytes)
	return activeUsers.GetUsername()
}
