package eventa

const BEFORE = -10

const AFTER = 10

type Hooks struct {
	Before []Callback
	After  []Callback
}
