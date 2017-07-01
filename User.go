package main

import (
	"strconv"
	"strings"
)

type User struct {
	name          string
	id            int
	lastRoom      *Room
	userVariables []*UserVariable
	session       Session
}

func (user *User) GetUserDataToString() string {
	userData := user.name + "^^" + strconv.Itoa(user.id) + "^^"

	for _, userVar := range user.userVariables {
		userData += userVar.key + "½" + userVar.value + "§"
	}
	userData = strings.TrimRight(userData, "#")
	return userData
}

func (user *User) AddUserVariables(variables []*UserVariable) {
	varIsExist := false

	for _, newVariable := range variables {
		for _, userVar := range user.userVariables {
			if newVariable.key == userVar.key {
				varIsExist = true
				userVar.value = newVariable.value
				break
			}
		}

		if !varIsExist {
			user.userVariables = append(user.userVariables, newVariable)
		}
	}
}
