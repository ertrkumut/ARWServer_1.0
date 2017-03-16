package main

 type ExtensionHandler func(*ARWServer, *User, ARWObject)

 type ExtensionRequest struct {
 	cmd     string
 	handler func(*ARWServer, *User, ARWObject)
 }
