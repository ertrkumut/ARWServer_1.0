package main

import "net"

type convert func(ARWObject)

type ARWEvent struct {
	eventName       string
	Handler         convert
	Private_Handler func(*ARWServer, net.Conn, ARWObject)
}
