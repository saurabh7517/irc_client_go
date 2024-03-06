package pkg

import (
	"bufio"
	"fmt"
	obj "irc_client/objects"
	"log"
	"net"
	"os"
)

func ProcessMessaging(conn net.Conn, token string) {
	//list all active users
	//create a go routine which will listen for new added active users
	var userList []string = getAllUsers(conn, token)
	if len(userList) > 0 {
		showActiveUsers(userList)
		fmt.Println("Enter the name of the user, you want to chat with...")
		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from console")
		}
		var found bool = false
		if len(userInput) > 0 {
			for _, user := range userList {
				if user == userInput {
					found = true
					break
				}
			}
		}
		if found {
			// send a message to the user on connection
			fmt.Print("Your message : >") //TODO own username
			message, _ := reader.ReadString('\n')

		} else {
			log.Println("User entered is not among active user list")
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
	var tokenMsg *obj.Message = CreateTokenMessage(token)
	var tokenMsgBytes []byte = EncodeMessage(tokenMsg)
	var activeUsers *obj.ActiveUser = SendRQGetActiveUsers(conn, tokenMsgBytes)
	return activeUsers.GetUsername()
}
