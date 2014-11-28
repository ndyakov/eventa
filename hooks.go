package eventa

const BEFORE = -10

const AFTER = 10

type Hooks struct {
	Before []Callback
	After  []Callback
}

func (h *Hooks) Initialize() {
	h.Before = make([]Callback, 1)
	h.After = make([]Callback, 1)
}
