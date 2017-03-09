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
