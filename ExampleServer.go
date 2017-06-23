package main

import "fmt"

var arwServer ARWServer

func main() {
	arwServer.Initialize()

	arwServer.AddExtensionHandler("FindRoom", FindRoomHandler)
	arwServer.ProcessEvents()
}

func FindRoomHandler(arwServer *ARWServer, user *User, obj ARWObject) {
	fmt.Println("===>> Find Room Request")
	roomTag, _ := obj.GetString("roomTag")

	findedRoom, err := arwServer.roomManager.SearchRoomWithTag(roomTag)

	if err != nil {
		// Create Room
		var roomSettings RoomSettings
		roomSettings.name = roomTag
		roomSettings.tag = roomTag
		roomSettings.cappacity = 4
		roomSettings.maxRoomVariableCount = 10
		roomSettings.InitializeMethod = RoomInitializeMethod

		newRoom := arwServer.roomManager.CreateRoom(roomSettings, arwServer)
		newRoom.AddUserToRoom(arwServer, user)
		return
	}

	findedRoom.AddUserToRoom(arwServer, user)
}

func RoomInitializeMethod(arwServer *ARWServer, room *Room) {
	fmt.Println("Room Initialize Success " + room.name)
	// room.AddExtensionHandler("IamReady", IamReadyHandler)
	room.AddExtensionHandler("VerticalUpdate", VerticalUpdateHandler)
	room.AddExtensionHandler("HorizontalUpdate", HorizontalUpdateHandler)
}

func VerticalUpdateHandler(arwServer *ARWServer, user *User, obj ARWObject) {
	obj.PutInt("userId", user.id)

	arwServer.roomManager.allRooms[0].SendRequestAllUserWithoutMe(*arwServer, obj, *user)
}

func HorizontalUpdateHandler(arwServer *ARWServer, user *User, obj ARWObject) {
	obj.PutInt("userId", user.id)

	arwServer.roomManager.allRooms[0].SendRequestAllUserWithoutMe(*arwServer, obj, *user)
}
