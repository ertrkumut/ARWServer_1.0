package main

import (
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn       net.Conn
	id         string
	createTime time.Time
}

//	--------------------- Getter Methods ---------------------

func (s *Session) GetConn() net.Conn {
	return s.conn
}

func (s *Session) GetConnString() string {
	return s.conn.RemoteAddr().String()
}

func (s *Session) GetId() string {
	return s.id
}

func (s *Session) GetCreatedTime() time.Time {
	return s.createTime
}

//	--------------------- Setter Methods ---------------------
func (s *Session) SetConn(conn net.Conn) {
	s.conn = conn
}

func (s *Session) SetId(id string) {
	s.id = id
}

func (s *Session) SetCreatedTime() {
	s.createTime = time.Now()
}

//	--------------------- ------ -------  ---------------------

func (s *Session) Init(conn *net.Conn) {
	s.SetCreatedTime()
	s.SetConn(*conn)
	fmt.Printf("Session Started : %s \n", s.GetConnString())
}
