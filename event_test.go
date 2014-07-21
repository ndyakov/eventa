package eventa_test

import (
	"./"
	"testing"
)

func TestEventStructShouldHaveEventID(t *testing.T) {
	var event eventa.EventID
	event = 10
	e := eventa.Event{Event: event}
	assert(t, e.Event == 10, "EventID is not set properly")
}
