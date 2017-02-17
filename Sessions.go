package main

import "net"

type Sessions interface {
	StartSession(network string)
}

type sessions struct {
	allSessions []Session
}

func (s sessions) StartSession(conn *net.Conn, sessionManager *sessions) {
	var ses Session

	ses.Init(conn, sessionManager)
}
