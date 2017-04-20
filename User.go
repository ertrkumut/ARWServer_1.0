package main

import (
	"fmt"
	"strconv"
)

type User struct {
	name     string
	id       int
	lastRoom Room
	session  Session
}

func (u *User) GetDataForOtherUser(user User) string {
	userData := u.name + "^^" + strconv.Itoa(u.id) + "^^"
	if u != &user {
		userData += "false" // IsMe true ^^ false
	} else {
		userData += "true"
	}

	return userData
}

func (u *User) ShutDownUser() { // kullanıcı serverdan düştükten sonra çalışan method
	fmt.Println("====> " + u.lastRoom.name)
	if u.lastRoom.name != "" {
		fmt.Println("Debug 3")
		u.lastRoom.RemoveUser(*u)
	}
}
