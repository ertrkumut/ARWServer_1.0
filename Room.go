package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type roomInitializeFunc func(arwServer *ARWServer, room *Room)

type Room struct {
	tag               string
	name              string
	password          string
	id                int
	cappacity         int
	maxVariableCount  int
	userList          []User
	roomVariables     []RoomVariable
	extensionHandlers []ExtensionRequest
	InitializeMethod  roomInitializeFunc
}

func (room *Room) Init(arwServer *ARWServer) {
	room.userList = make([]User, 0, room.cappacity)
	room.roomVariables = make([]RoomVariable, 0, room.maxVariableCount)

	if room.InitializeMethod != nil {
		room.InitializeMethod(arwServer, room)
	}
}

func (room *Room) AddUserToRoom(arwServer *ARWServer, u User) {
	room.userList = append(room.userList, u)

	var arwObj ARWObject

	arwObj.requestName = Join_Room // Odadaki tüm kullanıcı bilgileri odaya yeni giren elemana yollanır.
	arwObj.eventParams.PutString("RoomName", room.name)
	arwObj.eventParams.PutString("RoomTag", room.tag)
	arwObj.eventParams.PutInt("RoomId", room.id)
	arwObj.eventParams.PutInt("RoomCappacity", room.cappacity)

	usersData := ""
	for ii := 0; ii < len(room.userList); ii++ {
		if room.userList[ii].name != u.name {
			usersData += room.userList[ii].GetDataForOtherUser(u) + "''"
		}
	}

	usersData = strings.TrimRight(usersData, "''")

	arwObj.eventParams.PutString("Users", usersData)
	arwServer.SendRequestToUser(u, arwObj)

	var arwObjforTheOthers ARWObject
	arwObjforTheOthers.requestName = User_Enter_Room //Odaya yeni giren eleman'ın bilgileri odadaki diğer elemanlara yollanır
	arwObjforTheOthers.eventParams.PutString("RoomName", room.name)
	arwObjforTheOthers.eventParams.PutString("userName", u.name)
	arwObjforTheOthers.eventParams.PutInt("userId", u.id)
	arwObjforTheOthers.eventParams.PutString("isMe", "false")

	room.SendRequestAllUserWithoutMe(*arwServer, arwObjforTheOthers, u)
	fmt.Println("User join Room - User Name : ", u.name)
}

func (room *Room) IsFull() bool {
	if len(room.userList) < cap(room.userList) {
		return true
	}
	return false
}

func (room *Room) AddRoomVariables(variables []RoomVariable) {
	varIsExist := false
	for ii := 0; ii < len(variables); ii++ {
		for jj := 0; jj < len(room.roomVariables); jj++ {

			if variables[ii].key == room.roomVariables[jj].key {
				varIsExist = true
				room.roomVariables[jj].value = variables[ii].value
				break
			}
		}

		if !varIsExist {
			room.roomVariables = append(room.roomVariables, variables[ii])
		}
	}
}

func (room *Room) GetStringVariable(key string) (string, error) {
	for ii := 0; ii < len(room.roomVariables); ii++ {
		if room.roomVariables[ii].key == key {
			return room.roomVariables[ii].value, errors.New("")
		}
	}

	return "", errors.New("Variable does not exist")
}

func (room *Room) GetIntVariable(key string) (int, error) {

	for ii := 0; ii < len(room.roomVariables); ii++ {
		if room.roomVariables[ii].key == key {
			value, err := strconv.Atoi(room.roomVariables[ii].value)
			return value, err
		}
	}

	return 0, errors.New("Variable does not exist")
}

func (room *Room) GetFloatVariable(key string) (float64, error) {

	for ii := 0; ii < len(room.roomVariables); ii++ {
		if room.roomVariables[ii].key == key {
			value, err := strconv.ParseFloat(room.roomVariables[ii].value, 64)
			return value, err
		}
	}

	return 0, errors.New("Variable does not exist")
}

func (room *Room) GetBoolVariable(key string) (bool, error) {

	for ii := 0; ii < len(room.roomVariables); ii++ {
		if room.roomVariables[ii].key == key {
			value, err := strconv.ParseBool(room.roomVariables[ii].value)
			return value, err
		}
	}

	return false, errors.New("Variable does not exist")
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

func (room *Room) SendRequestAllUser(arwServer ARWServer, arwObj ARWObject) {

	if len(room.userList) == 0 {
		return
	}

	for ii := range room.userList {
		arwServer.SendRequestToUser(room.userList[ii], arwObj)
	}
}

func (room *Room) AddExtensionHandler(cmd string, handler ExtensionHandler) {
	if len(room.extensionHandlers) != 0 {
		for ii := 0; ii < len(room.extensionHandlers); ii++ {
			if cmd == room.extensionHandlers[ii].cmd {
				return
			}
		}
	}

	var newExtension *ExtensionRequest
	newExtension.cmd = cmd
	newExtension.handler = handler

	room.extensionHandlers = append(room.extensionHandlers, *newExtension)
}
