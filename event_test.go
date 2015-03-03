package eventa_test

import (
	"testing"

	"github.com/ndyakov/eventa"
)

func TestEventStructShouldHaveEventID(t *testing.T) {
	e := eventa.NewEvent("test::Event")
	assert(t, e.Name == "test::Event", "Event Name is not set properly")
}
