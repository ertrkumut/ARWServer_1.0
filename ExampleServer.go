package main

import "fmt"

var arwServer ARWServer

func main() {

	arwServer.Initialize()

	// arwServer.AddEventHandler(&arwServer.events.Login, Loginhandler)

	arwServer.AddExtensionHandler("Hello", HelloHandler)
	arwServer.AddEventHandler(&arwServer.events.Room_Create, RoomCreatehandler)
	arwServer.ProcessEvents()
}

func HelloHandler(arwServer *ARWServer, user *User, obj ARWObject) {
	fmt.Println(user.name, obj.GetString("Deneme"))
}

func RoomCreatehandler(obj ARWObject) {
	roomId := obj.eventParams.GetInt("RoomId")

	room, err := arwServer.roomManager.FindRoomWithID(roomId)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Room Create : ", roomId)

	room.InitializeMethod = RoomInitialize
}

func RoomInitialize(server *ARWServer, room *Room) {
	room.AddExtensionHandler("Hello", HelloHandler)
	fmt.Println(room.name, "intialize")
}
