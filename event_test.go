package eventa_test

import (
	"github.com/ndyakov/eventa"
	"testing"
)

func TestEventStructShouldHaveEventID(t *testing.T) {
	e := eventa.NewEvent(10)
	assert(t, e.Event == 10, "EventID is not set properly")
}
