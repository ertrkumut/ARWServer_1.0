package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ARWObject struct {
	requestName string
	dataList    map[string]string
	eventParams SpecialEventParam
}

func (arwObj *ARWObject) PutString(key string, value string) {

	if arwObj.dataList == nil {
		arwObj.dataList = make(map[string]string)
	}

	if arwObj.dataList[key] != "" {
		fmt.Println("The Key already exist...")
		return
	}

	arwObj.dataList[key] = value
}

func (arwObj *ARWObject) PutFloat(key string, value float64) {

	if arwObj.dataList == nil {
		arwObj.dataList = make(map[string]string)
	}

	if arwObj.dataList[key] != "" {
		fmt.Println("The Key already exist...")
		return
	}

	arwObj.dataList[key] = strconv.FormatFloat(value, 'f', -1, 64)
}

func (arwObj *ARWObject) PutInt(key string, value int) {

	if arwObj.dataList == nil {
		arwObj.dataList = make(map[string]string)
	}

	if arwObj.dataList[key] != "" {
		fmt.Println("The Key already exist...")
		return
	}

	arwObj.dataList[key] = strconv.Itoa(value)
}

func (arwObj *ARWObject) PutBool(key string, value bool) {

	if arwObj.dataList == nil {
		arwObj.dataList = make(map[string]string)
	}

	if arwObj.dataList[key] != "" {
		fmt.Println("The Key already exist...")
		return
	}

	arwObj.dataList[key] = strconv.FormatBool(value)
}

func (arwObj *ARWObject) GetString(key string) (string, error) {

	for k, v := range arwObj.dataList {
		if k == key {
			return v, nil
		}
	}
	return "", errors.New("Variable does not exist")
}

func (arwObj *ARWObject) GetFloat(key string) (float64, error) {
	for k, v := range arwObj.dataList {
		if k == key {
			value, convertErr := strconv.ParseFloat(v, 64)
			if convertErr != nil {
				return value, convertErr
			}
			return value, nil
		}
	}
	return 0, errors.New("Variable does not exist")
}

func (arwObj *ARWObject) GetInt(key string) (int, error) {

	for k, v := range arwObj.dataList {
		if k == key {
			value, convertErr := strconv.Atoi(v)
			if convertErr != nil {
				return value, convertErr
			}
			return value, nil
		}
	}
	return 0, errors.New("Variable does not exist")
}

func (arwObj *ARWObject) GetBool(key string) (bool, error) {

	for k, v := range arwObj.dataList {
		if k == key {
			value, convertErr := strconv.ParseBool(v)
			if convertErr != nil {
				return value, convertErr
			}
			return value, nil
		}
	}

	return false, errors.New("Variable does not exist")
}

func (arwObj *ARWObject) Compress() []byte {
	var data string

	data += arwObj.requestName + "~"

	for k, v := range arwObj.dataList {
		data += k + "#" + v + "_"
	}

	data = strings.TrimRight(data, "_")
	data += "~"
	data += arwObj.eventParams.Compress()
	data += "|"
	bytes := make([]byte, 1024)
	bytes = []byte(data)

	return bytes
}

func (arwObj *ARWObject) Extract(bytes []byte) {
	data := string(bytes)

	dataParts := strings.Split(data, "~")
	if len(dataParts) == 3 {
		reqName := dataParts[0]
		arwObj.requestName = reqName

		params := strings.Split(dataParts[1], "_")
		arwObj.dataList = make(map[string]string)

		for ii := 0; ii < len(params); ii++ {
			paramParts := strings.Split(params[ii], "#")

			if len(paramParts) == 2 {
				arwObj.dataList[paramParts[0]] = paramParts[1]
			}
		}

		arwObj.eventParams.Extract(dataParts[2])

	}
}

func (arwObj *ARWObject) GetUser(arwServer ARWServer) (*User, error) {
	id, _ := arwObj.eventParams.GetInt("userId")

	user, err := arwServer.userManager.FindUserWithId(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
