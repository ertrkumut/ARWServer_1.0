package main

import "errors"

type RoomManager struct {
	allRooms    []Room
	roomCounter int
}

func (roomManager *RoomManager) CreateRoom(settings RoomSettings, arwServer *ARWServer) *Room {

	var newRoom Room
	newRoom.InitializeMethod = settings.InitializeMethod

	newRoom.name = settings.name
	newRoom.password = settings.password
	newRoom.tag = settings.tag
	newRoom.cappacity = settings.cappacity
	newRoom.maxVariableCount = settings.maxRoomVariableCount

	newRoom.id = roomManager.roomCounter
	roomManager.roomCounter++

	roomManager.allRooms = append(roomManager.allRooms, newRoom)

	var roomCreateArwObj ARWObject
	roomCreateArwObj.requestName = Room_Create
	roomCreateArwObj.eventParams.PutString("RoomName", newRoom.name)
	roomCreateArwObj.eventParams.PutString("RoomTag", newRoom.tag)
	roomCreateArwObj.eventParams.PutInt("RoomId", newRoom.id)
	roomCreateArwObj.eventParams.PutInt("RoomCappacity", settings.cappacity)

	if arwServer.events.Room_Create.Handler != nil {
		arwServer.events.Room_Create.Handler(roomCreateArwObj)
	}

	arwServer.sessions.SendRequestToAllSessions(arwServer, roomCreateArwObj)

	newRoom.Init(arwServer)
	return &newRoom
}

func (roomManager *RoomManager) SearchRoomWithTag(roomTag string) (*Room, string) {
	var currentRoom *Room

	if len(roomManager.allRooms) == 0 {
		return currentRoom, "There was no room"
	}

	for ii := 0; ii < len(roomManager.allRooms); ii++ {
		if !roomManager.allRooms[ii].IsFull() {
			currentRoom = &roomManager.allRooms[ii]
			return currentRoom, ""
		}
	}

	return currentRoom, "There was no room"
}

func (roomManager *RoomManager) FindRoomWithID(id int) (Room, error) {

	for ii := 0; ii < len(roomManager.allRooms); ii++ {
		if roomManager.allRooms[ii].id == id {
			return roomManager.allRooms[ii], nil
		}
	}

	var r Room
	return r, errors.New("Room Not Found")
}
