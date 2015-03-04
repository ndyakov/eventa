package eventa

import "github.com/ndyakov/whatever"

type Callback func(*Listener, whatever.Params)

type callbacks struct {
	Permanent []Callback
	Once      []Callback
}

func newCallbacks() *callbacks {
	c := &callbacks{}
	c.initialize()
	return c
}

func (c *callbacks) initialize() {
	c.Permanent = make([]Callback, 1)
	c.Once = make([]Callback, 1)
}
