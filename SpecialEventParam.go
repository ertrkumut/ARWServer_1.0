package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SpecialEventParam struct {
	dataList map[string]string
}

func (evntParam *SpecialEventParam) PutString(key string, value string) {
	if evntParam.dataList == nil {
		evntParam.dataList = make(map[string]string)
	}

	if evntParam.dataList[key] != "" {
		fmt.Println("The key already exist")
		return
	}

	evntParam.dataList[key] = value
}

func (evntParam *SpecialEventParam) PutFloat(key string, value float64) {

	if evntParam.dataList == nil {
		evntParam.dataList = make(map[string]string)
	}

	if evntParam.dataList[key] != "" {
		fmt.Println("The key already exist")
		return
	}

	evntParam.dataList[key] = strconv.FormatFloat(value, 'f', -1, 64)
}

func (evntParam *SpecialEventParam) PutInt(key string, value int) {

	if evntParam.dataList == nil {
		evntParam.dataList = make(map[string]string)
	}

	if evntParam.dataList[key] != "" {
		fmt.Println("The key already exist")
		return
	}

	evntParam.dataList[key] = strconv.Itoa(value)
}

func (evntParam *SpecialEventParam) GetString(key string) (string, error) {

	for k, v := range evntParam.dataList {
		if k == key {
			return v, nil
		}
	}

	return "", errors.New("Variable does not exist")
}

func (evntParam *SpecialEventParam) GetFloat(key string) (float64, error) {

	for k, v := range evntParam.dataList {
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

func (evntParam *SpecialEventParam) GetInt(key string) (int, error) {

	for k, v := range evntParam.dataList {
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

func (evntParam *SpecialEventParam) Extract(data string) {

	dataParts := strings.Split(data, "_")
	evntParam.dataList = make(map[string]string)

	for ii := 0; ii < len(dataParts); ii++ {
		varParts := strings.Split(dataParts[ii], "#")
		if len(varParts) == 2 {
			evntParam.dataList[varParts[0]] = varParts[1]
		}
	}
}

func (evntParam *SpecialEventParam) Compress() string {
	var data string
	data = ""

	for k, v := range evntParam.dataList {
		data += k + "#" + v + "_"
	}

	data = strings.TrimRight(data, "_")
	return data
}
