package main

import "strconv"

type User struct {
	name         string
	id           int
	lastRoom     *Room
	userVariable []*UserVariable
	session      Session
}

func (user *User) GetDataForOtherUser(otherUser *User) string {
	userData := user.name + "^^" + strconv.Itoa(user.id) + "^^"
	if user != otherUser {
		userData += "false" // IsMe true ^^ false
	} else {
		userData += "true"
	}

	return userData
}

func (user *User) AddUserVariable(variables []*UserVariable) {
	varIsExist := false

	for _, newVariable := range variables {
		for _, userVar := range user.userVariable {
			if newVariable.key == userVar.key {
				varIsExist = true
				userVar.value = newVariable.value
				break
			}
		}

		if !varIsExist {
			user.userVariable = append(user.userVariable, newVariable)
		}
	}
}
