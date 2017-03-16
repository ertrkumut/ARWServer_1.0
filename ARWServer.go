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
	conn.Write(obj.Compress())
}

func (arw *ARWServer) SendRequestToUser(user User, obj ARWObject) {
	user.session.GetConn().Write(obj.Compress())
}

func (arw *ARWServer) SendExceptionToUser(user User, err string) {
	var obj ARWObject

	obj.requestName = Exception_Error
	obj.eventParams.PutString("error", err)

	arw.SendRequestToUser(user, obj)
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
		// timeoutDuration := 1 * time.Millisecond

		// conn.SetReadDeadline(time.Now().Add(timeoutDuration))
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
				arw.sessions.CloseSession(conn)
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

	var newExtension *ExtensionRequest
	newExtension.cmd = cmd
	newExtension.handler = handler

	arw.extensionHandlers = append(arw.extensionHandlers, *newExtension)
}
