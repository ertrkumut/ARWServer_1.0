package main

import "fmt"

func main() {
	var arwServer ARWServer
	arwServer.Initialize()

	arwServer.AddEventHandler(&arwServer.events.User_Login_Event, Loginhandler)
	arwServer.ProcessEvents()
}

func Loginhandler(arwObj ARWObject) {
	usrName := arwObj.evntParams.GetString("userName")

	fmt.Println("Manuel Login Handler : ", usrName)
}
