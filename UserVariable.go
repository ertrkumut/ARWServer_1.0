package main

import "strconv"

type UserVariable struct {
	key   string
	value string
}

func new_UserVariable_String(key string, value string) *UserVariable {
	var newVariable *UserVariable
	newVariable.key = key
	newVariable.value = value
	return newVariable
}

func new_UserVariable_Int(key string, value int) *UserVariable {
	var newVariable *UserVariable
	newVariable.key = key
	newVariable.value = string(value)
	return newVariable
}

func new_UserVariable_Float(key string, value float64) *UserVariable {
	var newVariable *UserVariable
	newVariable.key = key
	newVariable.value = strconv.FormatFloat(value, 'f', -1, 64)
	return newVariable
}

func new_UserVariable_Bool(key string, value bool) *UserVariable {
	var newVariable *UserVariable
	newVariable.key = key
	newVariable.value = strconv.FormatBool(value)
	return newVariable
}

func (variable *UserVariable) GetValueToString() string {
	return variable.value
}

func (variable *UserVariable) GetValueToFloat() float64 {
	value, _ := strconv.ParseFloat(variable.value, 64)
	return value
}

func (variable *UserVariable) GetValueToInt() int {
	value, _ := strconv.Atoi(variable.value)
	return value
}

func (variable *UserVariable) GetValueToBool() bool {
	value, _ := strconv.ParseBool(variable.value)
	return value
}
