package eventa

import "github.com/ndyakov/whatever"

type Event struct {
	Name   string
	Params whatever.Params
}

func NewEvent(eventName string) *Event {
	e := &Event{Name: eventName}
	e.Params = whatever.Params{}
	return e
}
