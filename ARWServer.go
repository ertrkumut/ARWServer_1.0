package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

type ARWServer struct {
	sessions    SessionManager
	userManager UserManager
	listenner   net.Listener
	events      ARWEvents
}

func (arw *ARWServer) SendRequestWithConn(conn net.Conn, obj ARWObject) {
	conn.Write(obj.Compress())
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
			P_ConnectionSuccess(arw, conn, arwObj)
		} else if arwObj.requestName == "LoginEvent" {
			P_LoginEvent(arw, conn, arwObj)
			fmt.Println(len(arw.userManager.allUsers))
		}
	}

}

func main() {
	var arwServer ARWServer
	arwServer.Initialize()

	arwServer.ProcessEvents()
}
