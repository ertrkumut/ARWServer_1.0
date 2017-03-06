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
	newArwObj.eventParams.PutString("isMe", "true")

	if arwServer.events.Login.Handler != nil {
		arwServer.events.Login.Handler(newArwObj)
	}

	arwServer.SendRequestWithConn(conn, newArwObj)
}

func P_JoinAnyRoom(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	roomTag := arwObj.eventParams.GetString("roomTag")

	r, e := arwServer.roomManager.FindRoom(roomTag) // Boş bir oda bulmak için odaları denetle

	if e != "" { // Eğer oda bulunamazsa odayı oluştur.
		r = arwServer.roomManager.CreateRoom(roomTag)
	}
	var currentUser User

	for ii := 0; ii < len(arwServer.userManager.allUsers); ii++ { // Paketin geldiği sessiondaki user bulunuyor.
		if conn == arwServer.userManager.allUsers[ii].session.GetConn() {
			currentUser = arwServer.userManager.allUsers[ii]
			break
		}
	}

	if currentUser.name != "" { // User odaya ekleniyor.
		r.AddUser(arwServer, currentUser)
	}
}
