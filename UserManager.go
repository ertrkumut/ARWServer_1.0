package main

import (
	"fmt"
	"net"
	"os"
)

type UserManager struct {
	allUsers []User
}

func (userManager *UserManager) CreateUser(userName string, conn net.Conn, arwServer *ARWServer) *User {
	var newUser User
	newUser.SetName(userName)

	var ses Session
	for ii := range arwServer.sessions.allSessions {
		currentSession := arwServer.sessions.allSessions[ii]
		if conn == currentSession.GetConn() {
			ses = currentSession
			newUser.SetSession(ses)
		}
	}

	if ses.conn == nil {
		fmt.Println("Session Not Found!! ")
		os.Exit(1)
	}

	newUser.SetId(len(arwServer.userManager.allUsers))
	arwServer.userManager.allUsers = append(arwServer.userManager.allUsers, newUser)
	return &newUser
}
