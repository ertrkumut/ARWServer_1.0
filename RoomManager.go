package main

type RoomManager struct {
	allRooms    []Room
	roomCounter int
}

func (roomManager *RoomManager) CreateRoom(name string, tag string, cappacity int, arwServer *ARWServer) *Room {

	var newRoom Room
	newRoom.userList = make([]User, 0, cappacity)
	newRoom.roomVariables = make([]RoomVariable, 0, 5)

	newRoom.name = name
	newRoom.tag = tag
	newRoom.id = roomManager.roomCounter
	roomManager.roomCounter++

	roomManager.allRooms = append(roomManager.allRooms, newRoom)

	var roomCreateArwObj ARWObject
	roomCreateArwObj.requestName = Room_Create
	roomCreateArwObj.eventParams.PutString("RoomName", newRoom.name)
	roomCreateArwObj.eventParams.PutString("RoomTag", newRoom.tag)
	roomCreateArwObj.eventParams.PutInt("RoomId", newRoom.id)
	roomCreateArwObj.eventParams.PutInt("RoomCappacity", cap(newRoom.userList))

	if arwServer.events.Room_Create.Handler != nil {
		arwServer.events.Room_Create.Handler(roomCreateArwObj)
	}

	arwServer.sessions.SendRequestToAllSessions(arwServer, roomCreateArwObj)

	return &newRoom
}

func (roomManager *RoomManager) FindRoom(roomTag string) (*Room, string) {
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
