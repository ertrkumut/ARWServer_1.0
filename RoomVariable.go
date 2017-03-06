package main

import "strconv"

type RoomVariable struct {
	key   string
	value string
}

func new_RoomVariable_String(key string, value string) *RoomVariable {
	var newVariable *RoomVariable
	newVariable.key = key
	newVariable.value = value
	return newVariable
}

func new_RoomVariable_Int(key string, value int) *RoomVariable {
	var newVariable *RoomVariable
	newVariable.key = key
	newVariable.value = string(value)
	return newVariable
}

func new_RoomVariable_Float(key string, value float64) *RoomVariable {
	var newVariable *RoomVariable
	newVariable.key = key
	newVariable.value = strconv.FormatFloat(value, 'f', -1, 64)
	return newVariable
}

func new_RoomVariable_Bool(key string, value bool) *RoomVariable {
	var newVariable *RoomVariable
	newVariable.key = key
	newVariable.value = strconv.FormatBool(value)
	return newVariable
}

func (variable *RoomVariable) GetValueToString() string {
	return variable.value
}

func (variable *RoomVariable) GetValueToFloat() float64 {
	value, _ := strconv.ParseFloat(variable.value, 64)
	return value
}

func (variable *RoomVariable) GetValueToInt() int {
	value, _ := strconv.Atoi(variable.value)
	return value
}

func (variable *RoomVariable) GetValueToBool() bool {
	value, _ := strconv.ParseBool(variable.value)
	return value
}
