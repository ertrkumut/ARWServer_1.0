package main

import (
	"fmt"
	"net"
)

func P_ConnectionSuccess(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
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

func P_JoinAnyRoom(arwServer *ARWServer, conn net.Conn, arwObj ARWObject) {
	roomTag := arwObj.eventParams.GetString("roomTag")

	r, e := arwServer.roomManager.SearchRoomWithTag(roomTag) // Boş bir oda bulmak için odaları denetle

	if e != "" { // Eğer oda bulunamazsa odayı oluştur.
		r = arwServer.roomManager.CreateRoom("game", "GameRoom", 4, arwServer)
	}

	currentUser, err := arwServer.userManager.FindUserWithConn(conn)

	if err != nil {
		fmt.Println("User not found exception := P_JoinAnyRoom")
	}

	if currentUser.name != "" { // User odaya ekleniyor.
		r.AddUserToRoom(arwServer, currentUser)
	}
}
