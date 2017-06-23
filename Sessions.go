package main

import (
	"fmt"
	"net"
)

type SessionManager struct {
	allSessions []Session
}

func (s *SessionManager) StartSession(conn *net.Conn) {
	var ses Session

	ses.Init(conn, s)
}

func (s *SessionManager) SendRequestToAllSessions(arwServer *ARWServer, obj ARWObject) {

	for ii := 0; ii < len(s.allSessions); ii++ {
		arwServer.SendRequestWithConn(s.allSessions[ii].GetConn(), obj)
	}
}

func (s *SessionManager) CloseSession(arw *ARWServer, conn net.Conn) {
	newSessionsArray := make([]Session, 0, len(s.allSessions)-1)
	newUsersArray := make([]*User, 0, len(arw.userManager.allUsers)-1)

	for ii := 0; ii < len(s.allSessions); ii++ {
		if s.allSessions[ii].GetConn() != conn {
			newSessionsArray = append(newSessionsArray, s.allSessions[ii])
		}
	}
	s.allSessions = newSessionsArray

	var deleteduser *User
	for ii := 0; ii < len(arw.userManager.allUsers); ii++ {
		if arw.userManager.allUsers[ii].session.GetConn() != conn {
			newUsersArray = append(newUsersArray, arw.userManager.allUsers[ii])
		} else {
			deleteduser = arw.userManager.allUsers[ii]
		}
	}

	fmt.Println("User Disconnected : " + deleteduser.name + " Room : " + deleteduser.lastRoom.name)
	arw.userManager.allUsers = newUsersArray

	conn.Close()
}
