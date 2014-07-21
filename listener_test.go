package eventa_test

import (
	"./"
	"testing"
)

func TestListener_Start(t *testing.T) {
	l := new(eventa.Listener)
	ok(t, l.Start(10))
	notok(t, l.Start(10))
}

func TestListener_RegisterIfNil(t *testing.T) {
	l := new(eventa.Listener)

	ok(t, l.RegisterIfNil(1, func(l *eventa.Listener, ed eventa.EventData) {}))
	notok(t, l.RegisterIfNil(1, func(l *eventa.Listener, ed eventa.EventData) {}))
}
