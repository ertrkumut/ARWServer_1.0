package main

import "errors"

type RoomManager struct {
	allRooms    []*Room
	roomCounter int
}

func (roomManager *RoomManager) CreateRoom(settings RoomSettings, arwServer *ARWServer) *Room {

	var newRoom *Room
	newRoom = new(Room)
	newRoom.InitializeMethod = settings.InitializeMethod

	newRoom.name = settings.name
	newRoom.password = settings.password
	newRoom.tag = settings.tag
	newRoom.cappacity = settings.cappacity
	newRoom.maxVariableCount = settings.maxRoomVariableCount

	newRoom.id = roomManager.roomCounter
	roomManager.roomCounter++

	newRoom.userList = make([]*User, 0, newRoom.cappacity)
	newRoom.roomVariables = make([]*RoomVariable, 0, newRoom.maxVariableCount)

	if newRoom.InitializeMethod != nil {
		newRoom.InitializeMethod(arwServer, newRoom)
	}

	roomManager.allRooms = append(roomManager.allRooms, newRoom)
	return newRoom
}

func (roomManager *RoomManager) SearchRoomWithTag(roomTag string) (*Room, error) {
	var currentRoom *Room

	for ii := 0; ii < len(roomManager.allRooms); ii++ {
		if roomManager.allRooms[ii].IsFull() == true {
			currentRoom = roomManager.allRooms[ii]
			return currentRoom, nil
		}
	}

	return currentRoom, errors.New("There was no room")
}

func (roomManager *RoomManager) FindRoomWithID(id int) (*Room, error) {

	for ii := 0; ii < len(roomManager.allRooms); ii++ {
		if roomManager.allRooms[ii].id == id {
			return roomManager.allRooms[ii], nil
		}
	}

	var r *Room
	return r, errors.New("Room Not Found")
}
