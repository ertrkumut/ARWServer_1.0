package main

type RoomManager struct {
	allRooms    []Room
	roomCounter int
}

func (roomManager *RoomManager) CreateRoom(name string, arwServer *ARWServer) *Room {

	var newRoom Room
	newRoom.userList = make([]User, 0, 4)
	newRoom.roomVariables = make([]RoomVariable, 0, 5)

	newRoom.name = name
	newRoom.id = roomManager.roomCounter
	roomManager.roomCounter++

	roomManager.allRooms = append(roomManager.allRooms, newRoom)

	var roomCreateArwObj ARWObject
	roomCreateArwObj.eventParams.PutInt("roomId", newRoom.id)

	if arwServer.events.Room_Create.Handler != nil {
		arwServer.events.Room_Create.Handler(roomCreateArwObj)
	}
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
