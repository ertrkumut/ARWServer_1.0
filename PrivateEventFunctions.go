package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func P_ConnectionSuccess(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	t := time.Now()
	s := t.Format(time.StampMilli)
	a := strings.Split(s, " ")

	realTime := a[2]

	arwObj.PutString("serverTime", realTime)
	arwServer.sessions.StartSession(&conn)
	arwObj.PutString("error", "")
	arwServer.SendRequestWithConn(conn, arwObj)
}

func P_LoginEvent(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	userName := arwObj.eventParams.GetString("userName")

	if arwServer.userManager.IsUserExist(userName) {
		var loginErrorObj ARWObject
		loginErrorObj.requestName = Login_Error
		loginErrorObj.PutString("error", "User already exist")

		arwServer.SendRequestWithConn(conn, loginErrorObj)
		return
	}

	newUser := arwServer.userManager.CreateUser(userName, conn, arwServer)
	fmt.Printf("User Login Server : %s id : %d Session : %s \n", newUser.name, newUser.id, newUser.session.GetConnString())

	var newArwObj ARWObject
	newArwObj.requestName = Login
	newArwObj.eventParams.PutString("userName", newUser.name)
	newArwObj.eventParams.PutInt("userId", newUser.id)
	newArwObj.eventParams.PutString("isMe", "true")

	if arwServer.events.Login.Handler != nil {
		arwServer.events.Login.Handler(newArwObj)
	}

	arwServer.SendRequestWithConn(conn, newArwObj)
}

func P_ExtensionResponse(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	if arwObj.eventParams.GetString("isRoomRequest") == "True" {
		currentRoom, err := arwServer.roomManager.FindRoomWithID(arwObj.eventParams.GetInt("roomId"))

		if err != nil {
			return
		}

		for ii := 0; ii < len(currentRoom.extensionHandlers); ii++ {
			if arwObj.eventParams.GetString("cmd") == currentRoom.extensionHandlers[ii].cmd {
				extension := currentRoom.extensionHandlers[ii]
				user, err := arwServer.userManager.FindUserWithConn(conn)

				if err == nil {
					extension.handler(arwServer, user, arwObj)
				}
			}
		}
	} else {
		for ii := 0; ii < len(arwServer.extensionHandlers); ii++ {
			if arwObj.eventParams.GetString("cmd") == arwServer.extensionHandlers[ii].cmd {
				extension := arwServer.extensionHandlers[ii]
				user, err := arwServer.userManager.FindUserWithConn(conn)

				if err == nil {
					extension.handler(arwServer, user, arwObj)
				}
			}
		}
	}
}

func P_Disconnection(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	fmt.Println("User disconnected!!")
	arwServer.sessions.CloseSession(arwServer, conn)
}
