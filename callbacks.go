package eventa

import "github.com/ndyakov/whatever"

type Callback func(*Listener, whatever.Params)

type Callbacks struct {
	Permanent []Callback
	Once      []Callback
}

func NewCallbacks() *Callbacks {
	c := &Callbacks{}
	c.Initialize()
	return c
}

func (c *Callbacks) Initialize() {
	c.Permanent = make([]Callback, 1)
	c.Once = make([]Callback, 1)
}
