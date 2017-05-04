package main

import "strconv"

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
