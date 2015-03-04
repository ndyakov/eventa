package eventa

import (
	"errors"

	"github.com/ndyakov/whatever"
)

type Listener struct {
	numEvents int
	incoming  chan *Event
	callbacks map[string]*Callbacks
	active    bool
}

func NewListener(numEvents int) *Listener {
	l := new(Listener)
	l.Start(numEvents)
	return l
}

func (l *Listener) Start(numEvents int) (err error) {
	if l.active {
		err = errors.New("The listener is already active!")
		return
	}

	if numEvents < 1 {
		numEvents = 1
	}

	l.numEvents = numEvents
	l.incoming = make(chan *Event, l.numEvents)
	l.callbacks = make(map[string]*Callbacks)
	l.active = true

	go l.listen()

	return
}

func (l *Listener) listen() {
	for {
		e := <-l.incoming

		if e.Name == "eventa::STOP" {
			return
		}

		l.runCallbacks(e, true)
	}
}

func (l *Listener) Stop() (err error) {
	if !l.active {
		err = errors.New("The listener is already inactive!")
		return
	}

	l.active = false

	l.Emit(&Event{Name: "eventa::STOP", Params: whatever.Params{}})

	return
}

func (l *Listener) Emit(event *Event) {
	l.initIfNeeded(event.Name)
	l.incoming <- event
}

func (l *Listener) On(eventName string, callback Callback) {
	l.initIfNeeded(eventName)
	l.callbacks[eventName].Permanent = append(l.callbacks[eventName].Permanent, callback)
}

func (l *Listener) OnceOn(eventName string, callback Callback) {
	l.initIfNeeded(eventName)
	l.callbacks[eventName].Once = append(l.callbacks[eventName].Once, callback)
}

func (l *Listener) initIfNeeded(eventName string) {
	if !l.active {
		l.Start(1)
	}

	if _, ok := l.callbacks[eventName]; !ok {
		l.callbacks[eventName] = NewCallbacks()
	}
}

func (l *Listener) ListenOn(in chan *Event) {
	l.Stop()
	l.incoming = in
	go l.listen()
}

func (l *Listener) runCallbacks(e *Event, concurrent bool) {
	for _, callback := range l.callbacks[e.Name].Permanent {
		if callback != nil {
			go callback(l, e.Params)
		}
	}

	for _, callback := range l.callbacks[e.Name].Once {
		if callback != nil {
			go callback(l, e.Params)
		}
	}

	l.callbacks[e.Name].Once = make([]Callback, 1)
}
