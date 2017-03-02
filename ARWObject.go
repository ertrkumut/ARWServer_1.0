package main

import (
	"strconv"
	"strings"
)

type ARWObject struct {
	requestName string
	dataList    []map[string]string
	eventParams SpecialEventParam
}

func (arwObj *ARWObject) PutString(key string, value string) {
	newField := make(map[string]string)
	newField[key] = value
	arwObj.dataList = append(arwObj.dataList, newField)
}

func (arwObj *ARWObject) PutFloat(key string, value float64) {
	newField := make(map[string]string)
	newField[key] = strconv.FormatFloat(value, 'f', -1, 64)
	arwObj.dataList = append(arwObj.dataList, newField)
}

func (arwObj *ARWObject) PutInt(key string, value int) {
	newField := make(map[string]string)
	newField[key] = strconv.Itoa(value)
	arwObj.dataList = append(arwObj.dataList, newField)
}

func (arwObj *ARWObject) GetString(key string) (value string) {
	for ii := 0; ii < len(arwObj.dataList); ii++ {
		c := arwObj.dataList[ii]
		for k, v := range c {
			if k == key {
				value = v
				return
			}
		}
	}
	return
}
func (arwObj *ARWObject) GetFloat(key string) (value float64) {
	for ii := 0; ii < len(arwObj.dataList); ii++ {
		c := arwObj.dataList[ii]
		for k, v := range c {
			if k == key {
				value, _ = strconv.ParseFloat(v, 64)
				return
			}
		}
	}
	return
}

func (arwObj *ARWObject) GetInt(key string) (value int) {
	for ii := 0; ii < len(arwObj.dataList); ii++ {
		c := arwObj.dataList[ii]
		for k, v := range c {
			if k == key {
				value, _ = strconv.Atoi(v)
				return
			}
		}
	}
	return
}

func (arwObj *ARWObject) Compress() []byte {
	var data string

	data += arwObj.requestName + "."
	for ii := 0; ii < len(arwObj.dataList); ii++ {
		p := arwObj.dataList[ii]
		for k, v := range p {
			data += k + "#" + v + "_"
		}
	}
	data = strings.TrimRight(data, "_")
	bytes := make([]byte, 1024)
	bytes = []byte(data)
	return bytes
}

func (arwObj *ARWObject) Extract(bytes []byte) {
	data := string(bytes)

	dataParts := strings.Split(data, ".")
	if len(dataParts) == 3 {
		reqName := dataParts[0]
		arwObj.requestName = reqName

		params := strings.Split(dataParts[1], "_")

		for ii := 0; ii < len(params); ii++ {
			paramParts := strings.Split(params[ii], "#")

			newMap := make(map[string]string)

			if len(paramParts) == 2 {
				newMap[paramParts[0]] = paramParts[1]
				arwObj.dataList = append(arwObj.dataList, newMap)
			}
		}

		arwObj.eventParams.Extract(dataParts[2])
	}
}
