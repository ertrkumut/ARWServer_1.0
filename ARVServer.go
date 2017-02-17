package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

const (
	SUCCESS = "Success"
)

var sess sessions

func main() {

	sess.allSessions = make([]Session, 0, 100)

	fmt.Println("Launching server... ")

	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		fmt.Println("Error listening...")
		os.Exit(1)
	}
	defer ln.Close()

	for {

		conn, acceptErr := ln.Accept()

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

		HandleRequests(requestBytes, &conn)
	}
}

func ConnectionEvent(conn *net.Conn) {
	clientConn := *conn
	sess.StartSession(conn, &sess)
	clientConn.Write([]byte("ConnectionSuccess"))
}

func HandleRequests(reqBytes []byte, conn *net.Conn) {
	reqBytes = bytes.Trim(reqBytes, "\x00")
	reqData := string(reqBytes)

	if reqData != "" {
		if reqData == SUCCESS {
			ConnectionEvent(conn)
		}
	}
}
