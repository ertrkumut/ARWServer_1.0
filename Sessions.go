package main

import "net"

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

func (s *SessionManager) CloseSession(conn net.Conn) {
	newArray := make([]Session, 0, len(s.allSessions)-1)

	for ii := 0; ii < len(s.allSessions); ii++ {
		if s.allSessions[ii].GetConn() != conn {
			newArray = append(newArray, s.allSessions[ii])
		}
	}
	s.allSessions = newArray
	conn.Close()
}
