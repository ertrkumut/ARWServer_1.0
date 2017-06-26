package main

import (
	"strconv"
	"strings"
)

type User struct {
	name         string
	id           int
	lastRoom     *Room
	userVariable []*UserVariable
	session      Session
}

func (user *User) GetUserDataToString() string {
	userData := user.name + "^^" + strconv.Itoa(user.id) + "^^"

	for _, userVar := range user.userVariable {
		userData += userVar.key + "½" + userVar.value + "§"
	}
	userData = strings.TrimRight(userData, "#")
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
