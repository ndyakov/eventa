package eventa_test

import (
	"testing"

	"github.com/ndyakov/eventa"
)

func TestListener_Start(t *testing.T) {
	l := new(eventa.Listener)
	ok(t, l.Start())
	notok(t, l.Start())
}
