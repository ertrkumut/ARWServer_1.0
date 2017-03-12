package main

import (
	"fmt"
	"net"
	"os"
)

type UserManager struct {
	allUsers    []User
	userCounter int
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

	newUser.SetId(userManager.userCounter)
	userManager.userCounter++

	arwServer.userManager.allUsers = append(arwServer.userManager.allUsers, newUser)
	return &newUser
}

func (userManager *UserManager) IsUserExist(userName string) bool {

	if len(userManager.allUsers) == 0 {
		return false
	}

	for ii := 0; ii < len(userManager.allUsers); ii++ {
		if userManager.allUsers[ii].name == userName {
			return true
		}
	}

	return false
}

func (userManager *UserManager) FindUserWithConn(arwServer ARWServer, conn net.Conn) (User, string) {
	var user User
	for ii := 0; ii < len(userManager.allUsers); ii++ {
		if conn.RemoteAddr() == userManager.allUsers[ii].session.GetConn().RemoteAddr() {
			user = userManager.allUsers[ii]
			return user, ""
		}
	}

	return user, "user found exception"
}
