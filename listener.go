package eventa

import (
	"errors"

	"github.com/ndyakov/whatever"
)

type Listener struct {
	incoming  chan *Event
	callbacks map[string]*callbacks
	active    bool
}

func NewListener() *Listener {
	l := &Listener{
		make(chan *Event, 2),
		make(map[string]*callbacks),
		false,
	}
	l.Start()
	return l
}

func (l *Listener) Start() (err error) {
	if l.active {
		err = errors.New("The listener is already active!")
		return
	}

	l.active = true
	go l.listen()

	return
}

func (l *Listener) listen() {
	for {
		e := <-l.incoming

		if e == nil || e.Name == "eventa::STOP" {
			return
		}

		l.runCallbacks(e, true)
	}
}

func (l *Listener) Emit(event *Event) {
	go func() {
		l.incoming <- event
	}()
}

func (l *Listener) On(eventName string, callback Callback) {
	l.initIfNeeded(eventName)
	l.callbacks[eventName].Permanent = append(l.callbacks[eventName].Permanent, callback)
}

func (l *Listener) OnceOn(eventName string, callback Callback) {
	l.initIfNeeded(eventName)
	l.callbacks[eventName].Once = append(l.callbacks[eventName].Once, callback)
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

func (l *Listener) runCallbacks(e *Event, concurrent bool) {
	l.initIfNeeded(e.Name)
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

func (l *Listener) initIfNeeded(eventName string) {
	if _, ok := l.callbacks[eventName]; !ok {
		l.callbacks[eventName] = newCallbacks()
	}
}
