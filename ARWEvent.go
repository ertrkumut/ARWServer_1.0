package main

import "net"

type convert func(ARWObject)

type ARWEvent struct {
	eventName       string
	Handler         func(ARWObject)
	Private_Handler func(*ARWServer, net.Conn, ARWObject)
}
