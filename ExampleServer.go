package main

import "fmt"

var arwServer ARWServer

func main() {

	arwServer.Initialize()

	arwServer.AddExtensionHandler("FindRoom", FindRoomHandler)
	arwServer.ProcessEvents()
}

func FindRoomHandler(arwServer *ARWServer, user *User, obj ARWObject) {
	fmt.Println("Find Room Request")
	roomTag := obj.GetString("roomTag")

	findedRoom, err := arwServer.roomManager.SearchRoomWithTag(roomTag)

	if err != "" {
		// Create Room
		var roomSettings RoomSettings
		roomSettings.name = roomTag
		roomSettings.tag = roomTag
		roomSettings.cappacity = 4
		roomSettings.maxRoomVariableCount = 10
		roomSettings.InitializeMethod = RoomInitializeMethod

		newRoom := arwServer.roomManager.CreateRoom(roomSettings, arwServer)
		newRoom.AddUserToRoom(arwServer, *user)
		return
	}

	findedRoom.AddUserToRoom(arwServer, *user)
}

func RoomInitializeMethod(arwServer *ARWServer, room *Room) {
	fmt.Println("Room Initialize Success " + room.name)
	// room.AddExtensionHandler("IamReady", IamReadyHandler)
}

func IamReadyHandler(arwServer *ARWServer, user *User, obj ARWObject) {
	fmt.Println("I am ready " + user.name)
}
