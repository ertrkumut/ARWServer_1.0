package main

type extensionHandler func(*ARWServer, *User, ARWObject)

type ExtensionRequest struct {
	cmd     string
	handler func(*ARWServer, *User, ARWObject)
}
