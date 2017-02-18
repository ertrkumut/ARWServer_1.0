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

func (arwObj *ARWObject) PrepareJson() []byte {
	var data string
	data += "{"
	data += "\"RequestName\"" + ":\"" + arwObj.requestName + "\""
	data += ",\"Param" + "\":["
	for ii := 0; ii < len(arwObj.param); ii++ {
		p := arwObj.param[ii]

		for k, v := range p {
			data += "{\"" + k + "\": \"" + v + "\"},"
		}
	}
	data = strings.TrimRight(data, ",")
	data += "]}"

	return []byte(data)
}

func (arwObject *ARWObject) JsonToStruct(reqData string) {

}
