package main

type RoomManager struct {
	allRooms    []Room
	roomCounter int
}

func (roomManager *RoomManager) CreateRoom(name string) Room {
	var newRoom Room

	newRoom.name = name
	newRoom.id = roomManager.roomCounter
	roomManager.roomCounter++

	roomManager.allRooms = append(roomManager.allRooms, newRoom)
	return newRoom
}
