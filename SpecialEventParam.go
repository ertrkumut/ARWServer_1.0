package main

import "strings"

type SpecialEventParam struct {
	dataList []map[string]string
}

func (evntParam *SpecialEventParam) PutString(key string, value string) {
	newField := make(map[string]string)
	newField[key] = value
	evntParam.dataList = append(evntParam.dataList, newField)
}

func (evntParam *SpecialEventParam) GetString(key string) (value string) {
	for ii := 0; ii < len(evntParam.dataList); ii++ {
		c := evntParam.dataList[ii]
		for k, v := range c {
			if k == key {
				value = v
				return
			}
		}
	}

	return
}

func (evntParam *SpecialEventParam) Extract(data string) {

	dataParts := strings.Split(data, "_")

	for ii := 0; ii < len(dataParts); ii++ {
		varParts := strings.Split(dataParts[ii], "#")
		if len(varParts) == 2 {
			newMap := make(map[string]string)
			newMap[varParts[0]] = varParts[1]
			evntParam.dataList = append(evntParam.dataList, newMap)
		}
	}
}

func (evntParam *SpecialEventParam) Compress() []byte {
	var data string
	data = ""

	for ii := 0; ii < len(evntParam.dataList); ii++ {
		p := evntParam.dataList[ii]
		for k, v := range p {
			data += k + "#" + v + "_"
		}
	}
	data = strings.TrimRight(data, "_")
	bytes := make([]byte, 1024)
	bytes = []byte(data)
	return bytes
}
