package main

type convert func(ARWObject)

type ARWEvent struct {
	Handler func(ARWObject)
}
