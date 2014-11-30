package eventa

import (
	"errors"
)

const STOP = -1

type Callback func(*Listener, map[string]interface{})

type Listener struct {
	numEvents int
	incoming  chan *Event
	callbacks map[int]Callback
	hooks     map[int]*Hooks
	active    bool
}

func NewListener(numEvents int) *Listener {
	l := new(Listener)
	l.Start(numEvents)
	return l
}

func (l *Listener) Start(NumEvents int) (err error) {
	if l.active {
		err = errors.New("The listener is already active!")
		return
	}

	if NumEvents < 1 {
		NumEvents = 1
	}

	l.numEvents = NumEvents
	l.incoming = make(chan *Event, l.numEvents)
	l.callbacks = make(map[int]Callback)
	l.hooks = make(map[int]*Hooks)
	l.active = true

	l.Register(STOP, func(*Listener, map[string]interface{}) {})

	go l.listen()

	return
}

func (l *Listener) Stop() (err error) {
	if !l.active {
		err = errors.New("The listener is already inactive!")
		return
	}

	l.active = false

	l.Emit(&Event{Event: STOP, Data: map[string]interface{}{}})

	return
}

func (l *Listener) ListenOn(in chan *Event) {
	l.incoming = in
	l.Stop()
	go l.listen()
}

func (l *Listener) Register(event int, callback Callback) {
	if !l.active {
		l.Start(1)
	}

	if _, ok := l.hooks[event]; !ok {
		l.hooks[event] = new(Hooks)
		l.hooks[event].Initialize()
	}

	l.callbacks[event] = callback
}

func (l *Listener) RegisterIfNil(event int, callback Callback) (err error) {
	if l.callbacks[event] == nil {
		l.Register(event, callback)
		return
	}

	return errors.New("This event is already set!")
}

func (l *Listener) Registered(event int) bool {
	if l.callbacks[event] == nil {
		return false
	}
	return true
}

func (l *Listener) RegisterHook(t int, event int, callback Callback) {
	if t == BEFORE {
		l.hooks[event].Before = append(l.hooks[event].Before, callback)
	} else if t == AFTER {
		l.hooks[event].After = append(l.hooks[event].After, callback)
	}
}

func (l *Listener) Emit(event *Event) {
	l.incoming <- event
}

func (l *Listener) listen() {
	for {
		e := <-l.incoming

		if e.Event == STOP {
			return
		}

		l.runBeforeHooks(e)

		if e.Concurrent {
			go l.callbacks[e.Event](l, e.Data)
		} else {
			l.callbacks[e.Event](l, e.Data)
		}

		l.runAfterHooks(e)
	}
}

func (l *Listener) runBeforeHooks(e *Event) {
	l.runHooks(l.hooks[e.Event].Before, e)
}

func (l *Listener) runAfterHooks(e *Event) {
	l.runHooks(l.hooks[e.Event].After, e)
}

func (l *Listener) runHooks(hooks []Callback, e *Event) {
	for _, hook := range hooks {
		if hook != nil {
			if e.Concurrent {
				go hook(l, e.Data)
			} else {
				hook(l, e.Data)
			}
		}
	}
}
