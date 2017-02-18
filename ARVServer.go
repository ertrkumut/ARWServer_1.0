package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SUCCESS      = "Success"
	LOGINREQUEST = "LoginRequest"
)

type ARWServer struct {
	sessions  SessionManager
	listenner net.Listener
}

func (arw *ARWServer) Send(arwObject *ARWObject, user *User) (err error) {

	return nil
}

func (arw *ARWServer) Initialize() (err error) {

	fmt.Println("Launching server... ")

	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		fmt.Println("Error listening...")
		os.Exit(1)
	}

	arw.listenner = ln
	return err
}

func (arw *ARWServer) ProcessEvents() {
	conn, acceptErr := arw.listenner.Accept()

	if acceptErr != nil {
		fmt.Println("Error Accepting...")
		os.Exit(1)
	}

	requestBytes := make([]byte, 1024)

	_, err := conn.Read(requestBytes)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

}

func main() {

}
