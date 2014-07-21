package eventa

const Before = -10

const After = 10

type Hooks struct {
	Before []Callback
	After  []Callback
}

func (h *Hooks) Initialize() {
	h.Before = make([]Callback, 1)
	h.After = make([]Callback, 1)
}
