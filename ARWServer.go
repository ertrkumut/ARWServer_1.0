package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

type ARWServer struct {
	sessions          SessionManager
	userManager       UserManager
	listenner         net.Listener
	events            ARWEvents
	roomManager       RoomManager
	extensionHandlers []ExtensionRequest
}

func (arw *ARWServer) SendRequestWithConn(conn net.Conn, obj ARWObject) {
	go conn.Write(obj.Compress())
}

func (arw *ARWServer) SendRequestToUser(user User, obj ARWObject) {
	arw.SendRequestWithConn(user.session.GetConn(), obj)
}

func (arw *ARWServer) SendExceptionToUser(user User, err string) {
	var obj ARWObject

	obj.requestName = Exception_Error
	obj.eventParams.PutString("error", err)

	arw.SendRequestToUser(user, obj)
}

func (arw *ARWServer) SendExtensionRequestToUser(user User, cmd string, obj ARWObject) {
	obj.requestName = Extension_Response
	obj.eventParams.PutString("cmd", cmd)
	arw.SendRequestToUser(user, obj)
}

func (arw *ARWServer) SendExtensionRequestToUsers(users []User, cmd string, obj ARWObject) {
	obj.requestName = Extension_Response
	obj.eventParams.PutString("cmd", cmd)
	for ii := 0; ii < len(users); ii++ {
		arw.SendRequestToUser(users[ii], obj)
	}
}

func (arw *ARWServer) Initialize() {
	arw.events.Initialize()

	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		fmt.Println("Error listening...")
	}
	fmt.Print("Initialize Success... \n\n")

	arw.listenner = ln
}

func (arw *ARWServer) ProcessEvents() {
	for {
		conn, acceptErr := arw.listenner.Accept()

		if acceptErr != nil {
			fmt.Println("Error Accepting :", acceptErr)
		}

		go arw.HandleRequests(conn)
	}
}

func (arw *ARWServer) HandleRequests(conn net.Conn) {
	defer conn.Close()
	for {
		requestBytes := make([]byte, 1024)

		_, err := conn.Read(requestBytes)

		if err != nil {
			if err != io.EOF {
				println("Read to server failed:", err.Error())
			} else {
				arw.sessions.CloseSession(arw, conn)
				return
			}
		}

		requestBytes = bytes.Trim(requestBytes, "\x00")

		var arwObj ARWObject
		arwObj.Extract(requestBytes)

		for ii := 0; ii < len(arw.events.allEvents); ii++ {
			currentEvent := arw.events.allEvents[ii]
			if currentEvent.eventName == arwObj.requestName {
				currentEvent.Private_Handler(arw, conn, arwObj)
				break
			}
		}
	}
}

func (arw *ARWServer) AddEventHandler(event *ARWEvent, handler convert) {
	event.Handler = handler
}

func (arw *ARWServer) AddExtensionHandler(cmd string, handler ExtensionHandler) {
	if len(arw.extensionHandlers) != 0 {
		for ii := 0; ii < len(arw.extensionHandlers); ii++ {
			if cmd == arw.extensionHandlers[ii].cmd {
				return
			}
		}
	}

	var newExtension ExtensionRequest
	newExtension.cmd = cmd
	newExtension.handler = handler

	arw.extensionHandlers = append(arw.extensionHandlers, newExtension)
}
