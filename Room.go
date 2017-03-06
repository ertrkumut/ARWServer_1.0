package main

import "fmt"

type Room struct {
	tag      string
	name     string
	id       int
	userList []User
}

func (room *Room) AddUser(arwServer *ARWServer, u User) {
	room.userList = append(room.userList, u)

	var arwObj ARWObject

	arwObj.requestName = Join_Room
	arwObj.eventParams.PutString("roomName", room.name)
	arwObj.eventParams.PutString("roomTag", room.tag)

	arwServer.SendRequestToUser(u, arwObj)
	fmt.Println("User join Room - User Name : ", u.name)
}
