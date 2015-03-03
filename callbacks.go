package eventa

import "github.com/ndyakov/whatever"

type Callback func(*Listener, whatever.Params)

type Callbacks struct {
	Concurrent []Callback
	Sequential []Callback
	Once       []Callback
}

func NewCallbacks() *Callbacks {
	c := &Callbacks{}
	c.Initialize()
	return c
}

func (c *Callbacks) Initialize() {
	c.Concurrent = make([]Callback, 1)
	c.Sequential = make([]Callback, 1)
	c.Once = make([]Callback, 1)
}
