package main

import "net"

type convert func(net.Conn)

type ARWEvent struct {
	Handler func(net.Conn)
}
