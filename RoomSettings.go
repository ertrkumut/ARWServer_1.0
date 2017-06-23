package main

type RoomSettings struct {
	tag                  string
	name                 string
	password             string
	cappacity            int
	maxRoomVariableCount int
	variables            []RoomVariable
	InitializeMethod     roomInitializeFunc
}
