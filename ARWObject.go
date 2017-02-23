package main

import "strings"

type ARWObject struct {
	requestName string
	param       []map[string]string
}

func (arwObj *ARWObject) PutString(key string, value string) {
	newField := make(map[string]string)
	newField[key] = value
	arwObj.param = append(arwObj.param, newField)
}

func (arwObj *ARWObject) GetString(key string) (value string) {
	for ii := 0; ii < len(arwObj.param); ii++ {
		c := arwObj.param[ii]
		for k, v := range c {
			if k == key {
				value = v
				return
			}
		}
	}
	return
}

func (arwObj *ARWObject) CompressToARWObject() []byte {
	var data string

	data += arwObj.requestName + "."
	for ii := 0; ii < len(arwObj.param); ii++ {
		p := arwObj.param[ii]
		for k, v := range p {
			data += k + "#" + v + "_"
		}
	}
	data = strings.TrimRight(data, "_")
	bytes := make([]byte, 1024)
	bytes = []byte(data)
	return bytes
}

func (arwObj *ARWObject) ExtractToARWObject(bytes []byte) {
	data := string(bytes)

	dataParts := strings.Split(data, ".")
	reqName := dataParts[0]
	arwObj.requestName = reqName

	params := strings.Split(dataParts[1], "_")

	for ii := 0; ii < len(params); ii++ {
		paramParts := strings.Split(params[ii], "#")

		newMap := make(map[string]string)

		if len(paramParts) == 2 {
			newMap[paramParts[0]] = paramParts[1]
			arwObj.param = append(arwObj.param, newMap)
		}
	}
}
