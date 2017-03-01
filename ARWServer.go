package main

import (
	"bytes"
	"fmt"
	"io"
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
		// timeoutDuration := 1 * time.Millisecond

		// conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		if acceptErr != nil {
			fmt.Println("Error Accepting :", acceptErr)
			os.Exit(1)
		}

		go arw.HandleRequests(conn)
	}
}

func (arw *ARWServer) PrivateConnection(conn net.Conn) {
	arw.sessions.StartSession(&conn)
}

func (arw *ARWServer) HandleRequests(conn net.Conn) {
	defer conn.Close()
	for {
		requestBytes := make([]byte, 1024)

		_, err := conn.Read(requestBytes)
		if err != nil {
			if err != io.EOF {
				println("Read to server failed:", err.Error())
				os.Exit(1)
			} else {
				// println("EOF Fail")
			}
		}

		requestBytes = bytes.Trim(requestBytes, "\x00")

		var arwObj ARWObject
		arwObj.Extract(requestBytes)

		if arwObj.requestName == "ConnectionSuccess" {
			fmt.Println("Connection Success")
			conn.Write(arwObj.Compress())
		} else if arwObj.requestName == "LoginEvent" {
			fmt.Println("Login Event")
		}
	}

}

func main() {
	var arwServer ARWServer
	arwServer.Initialize()

	arwServer.ProcessEvents()
}
