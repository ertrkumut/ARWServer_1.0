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

	realTime := a[3]

	arwObj.PutString("serverTime", realTime)
	arwServer.sessionManager.StartSession(&conn)
	arwObj.PutString("error", "")
	arwServer.SendRequestWithConn(conn, arwObj)
}

func P_LoginEvent(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	userName, _ := arwObj.eventParams.GetString("userName")

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

	newArwObj.eventParams.PutString("userVariables", newUser.CompressUserVariables())
	arwServer.SendRequestWithConn(conn, newArwObj)
}

func P_ExtensionResponse(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	if res, err := arwObj.eventParams.GetString("isRoomRequest"); (err == nil) && (res == "True") {
		roomId, _ := arwObj.eventParams.GetInt("roomId")
		currentRoom, err := arwServer.roomManager.FindRoomWithID(roomId)

		if err != nil {
			return
		}

		for ii := 0; ii < len(currentRoom.extensionHandlers); ii++ {
			if res, err := arwObj.eventParams.GetString("cmd"); (err == nil) && (res == currentRoom.extensionHandlers[ii].cmd) {
				extension := currentRoom.extensionHandlers[ii]
				user, err := arwServer.userManager.FindUserWithConn(conn)

				if err == nil {
					extension.handler(arwServer, user, arwObj)
				}
			}
		}
	} else {
		for ii := 0; ii < len(arwServer.extensionHandlers); ii++ {
			if res, err := arwObj.eventParams.GetString("cmd"); (err == nil) && (res == arwServer.extensionHandlers[ii].cmd) {
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
	arwServer.sessionManager.CloseSession(arwServer, conn)
}
