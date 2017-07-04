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

	userData += user.CompressUserVariables()

	return userData
}

func (user *User) CompressUserVariables() string {
	userData := ""
	for _, userVar := range user.userVariables {
		userData += userVar.key + "½" + userVar.value + "§"
	}
	userData = strings.TrimRight(userData, "#")

	return userData
}

func (user *User) VariableIsExist(key string) (bool, *UserVariable) {
	variableIsExist := false

	for _, userVar := range user.userVariables {
		if userVar.key == key {
			variableIsExist = true
			return variableIsExist, userVar
		}
	}
	return variableIsExist, nil
}

func (user *User) AddUserVariables(variables []*UserVariable) {
	varIsExist := false

	for _, newVariable := range variables {
		for _, userVar := range user.userVariables {
			if newVariable.key == userVar.key && newVariable.value != userVar.value {
				varIsExist = true
				userVar.value = newVariable.value // Send User Variable Update func.
				break
			}
		}

		if !varIsExist {
			user.userVariables = append(user.userVariables, newVariable)
		}
	}
}

func (user *User) AddStringUserVariable(key string, value string) {
	varIsExist, existVar := user.VariableIsExist(key)

	if !varIsExist {
		var newVariable *UserVariable
		newVariable = new(UserVariable)

		newVariable.key = key
		newVariable.value = value

		user.userVariables = append(user.userVariables, newVariable)
	} else {
		if existVar.GetValueToString() != value {
			existVar.value = value
		}
	}
}
