package eventa_test

import (
	"github.com/ndyakov/eventa"
	"testing"
)

func TestHooks_Initialize(t *testing.T) {
	h := new(eventa.Hooks)
	assert(t, h.Before == nil, "Before is not nil before Initialize")
	assert(t, h.After == nil, "After is not nil before Initialize")

	h.Initialize()
	assert(t, h.Before != nil, "Before is nil after Initialize")
	assert(t, h.After != nil, "After is nil after Initialize")
}

func TestHooks_Before(t *testing.T) {
	equals(t, -10, eventa.BEFORE)
}

func TestHooks_After(t *testing.T) {
	equals(t, 10, eventa.AFTER)
}
