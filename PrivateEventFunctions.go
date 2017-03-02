package main

import (
	"fmt"
	"net"
)

func P_ConnectionSuccess(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	arwServer.sessions.StartSession(&conn)

	arwServer.SendRequestWithConn(conn, arwObj)
}

func P_LoginEvent(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	userName := arwObj.eventParams.GetString("userName")

	newUser := arwServer.userManager.CreateUser(userName, conn, arwServer)
	fmt.Printf("User Login Server : %s id : %d Session : %s \n", newUser.name, newUser.id, newUser.session.GetConnString())

	var newArwObj ARWObject
	newArwObj.requestName = Login
	newArwObj.eventParams.PutString("userName", newUser.name)
	newArwObj.eventParams.PutInt("userId", newUser.id)

	if arwServer.events.Login.Handler != nil {
		arwServer.events.Login.Handler(newArwObj)
	}

	arwServer.SendRequestWithConn(conn, newArwObj)
}
