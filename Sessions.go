package main

import "net"

type SessionManager struct {
	allSessions []Session
}

func (s SessionManager) StartSession(conn *net.Conn, sessionManager *SessionManager) {
	var ses Session

	ses.Init(conn, sessionManager)
}
