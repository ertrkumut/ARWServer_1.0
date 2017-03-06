package main

import "fmt"

type Room struct {
	tag           string
	name          string
	id            int
	userList      []User
	roomVariables []RoomVariable
}

func (room *Room) AddUser(arwServer *ARWServer, u User) {
	room.userList = append(room.userList, u)

	var arwObj ARWObject

	arwObj.requestName = Join_Room
	arwObj.eventParams.PutString("roomName", room.name)
	arwObj.eventParams.PutString("roomTag", room.tag)

	arwServer.SendRequestToUser(u, arwObj)

	var arwObjforTheOthers ARWObject
	arwObjforTheOthers.requestName = User_Enter_Room
	arwObjforTheOthers.eventParams.PutString("userName", u.name)
	arwObjforTheOthers.eventParams.PutInt("userId", u.id)
	arwObjforTheOthers.eventParams.PutString("isMe", "false")

	room.SendRequestAllUserWithoutMe(*arwServer, arwObjforTheOthers, u)
	fmt.Println("User join Room - User Name : ", u.name)
}

func (room *Room) IsFull() bool {
	if len(room.userList) < 4 {
		return true
	}
	return false
}

func (room *Room) SendRequestAllUserWithoutMe(arwServer ARWServer, arwObj ARWObject, user User) {

	if len(room.userList) == 0 {
		return
	}

	for ii := range room.userList {
		if room.userList[ii].id != user.id {
			arwServer.SendRequestToUser(room.userList[ii], arwObj)
		}
	}
}
