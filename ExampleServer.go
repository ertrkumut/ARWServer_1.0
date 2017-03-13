package main

func main() {
	var arwServer ARWServer
	arwServer.Initialize()

	// arwServer.AddEventHandler(&arwServer.events.Login, Loginhandler)
	arwServer.ProcessEvents()
}
