package eventa_test

import (
	"testing"

	"github.com/ndyakov/eventa"
)

func TestListener_Start(t *testing.T) {
	l := new(eventa.Listener)
	ok(t, l.Start(10))
	notok(t, l.Start(10))
}
