package main

import (
	"strconv"
	"strings"
)

type SpecialEventParam struct {
	dataList []map[string]string
}

func (evntParam *SpecialEventParam) PutString(key string, value string) {
	newField := make(map[string]string)
	newField[key] = value
	evntParam.dataList = append(evntParam.dataList, newField)
}

func (evntParam *SpecialEventParam) PutFloat(key string, value float64) {
	newField := make(map[string]string)
	newField[key] = strconv.FormatFloat(value, 'f', -1, 64)
	evntParam.dataList = append(evntParam.dataList, newField)
}

func (evntParam *SpecialEventParam) PutInt(key string, value int) {
	newField := make(map[string]string)
	newField[key] = strconv.Itoa(value)
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

func (evntParam *SpecialEventParam) GetFloat(key string) (value float64) {
	for ii := 0; ii < len(evntParam.dataList); ii++ {
		c := evntParam.dataList[ii]
		for k, v := range c {
			if k == key {
				value, _ = strconv.ParseFloat(v, 64)
				return
			}
		}
	}
	return
}

func (evntParam *SpecialEventParam) GetInt(key string) (value int) {
	for ii := 0; ii < len(evntParam.dataList); ii++ {
		c := evntParam.dataList[ii]
		for k, v := range c {
			if k == key {
				value, _ = strconv.Atoi(v)
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

func (evntParam *SpecialEventParam) Compress() string {
	var data string
	data = ""

	for ii := 0; ii < len(evntParam.dataList); ii++ {
		p := evntParam.dataList[ii]
		for k, v := range p {
			data += k + "#" + v + "_"
		}
	}
	data = strings.TrimRight(data, "_")
	return data
}
