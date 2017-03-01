package main

import (
	"fmt"
	"net"
)

func P_ConnectionSuccess(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	arwServer.sessions.StartSession(&conn)

	conn.Write(arwObj.Compress())
}

func P_LoginEvent(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	userName := arwObj.evntParams.GetString("userName")

	newUser := arwServer.userManager.CreateUser(userName, conn, arwServer)
	fmt.Printf("User Login Server : %s id : %d Session : %s \n", newUser.name, newUser.id, newUser.session.GetConnString())

	if arwServer.events.User_Login_Event.Handler != nil {
		arwServer.events.User_Login_Event.Handler(arwObj)
	}
}
