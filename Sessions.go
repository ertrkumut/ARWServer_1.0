package main

import "net"

type SessionManager struct {
	allSessions []Session
}

func (s *SessionManager) StartSession(conn *net.Conn) {
	var ses Session

	ses.Init(conn, s)
}
