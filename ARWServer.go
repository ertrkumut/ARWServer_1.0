package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

type ARWServer struct {
	sessions  SessionManager
	listenner net.Listener
	events    ARWEvents
}

func (arw *ARWServer) SendResponceToUser(arwObject *ARWObject, user *User) (err error) {

	return nil
}

func (arw *ARWServer) Initialize() {

	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		fmt.Println("Error listening...")
		os.Exit(1)
	}
	fmt.Print("Initialize Success... \n\n")

	arw.listenner = ln
}

func (arw *ARWServer) ProcessEvents() {
	for {
		conn, acceptErr := arw.listenner.Accept()

		if acceptErr != nil {
			fmt.Println("Error Accepting :", acceptErr)
			os.Exit(1)
		}

		requestBytes := make([]byte, 1024)

		_, err := conn.Read(requestBytes)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		requestBytes = bytes.Trim(requestBytes, "\x00")
		arw.HandleRequests(requestBytes, conn)
	}
}

func (arw *ARWServer) PrivateConnection(conn net.Conn) {
	arw.sessions.StartSession(&conn)
	conn.Write([]byte("ConnectionSuccess"))
}

func (arw *ARWServer) HandleRequests(req []byte, conn net.Conn) {
	var arwObj ARWObject
	arwObj.ExtractToARWObject(req)

	if arwObj.requestName == "ConnectionSuccess" {
		fmt.Println("Connection Success")
		conn.Write(arwObj.CompressToARWObject())
	}
}

func main() {
	var arwServer ARWServer
	arwServer.Initialize()

	arwServer.ProcessEvents()
}
