package main

import "fmt"

func main() {
	var arwServer ARWServer
	arwServer.Initialize()

	arwServer.AddEventHandler(&arwServer.events.Login, Loginhandler)
	arwServer.ProcessEvents()
}

func Loginhandler(arwObj ARWObject) {
	userName := arwObj.eventParams.GetString("userName")
	userId := arwObj.eventParams.GetInt("userId")
	fmt.Printf("Manuel Login Handler : %s Id : %d \n", userName, userId)
}
