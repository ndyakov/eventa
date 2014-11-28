package eventa

import (
	"errors"
)

const stopListener = -1

type Callback func(*Listener, EventData)

type Listener struct {
	numEvents int
	incoming  chan Event
	callbacks map[EventID]Callback
	hooks     map[EventID]*Hooks
	active    bool
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
	l.incoming = make(chan Event, l.numEvents)
	l.callbacks = make(map[EventID]Callback)
	l.hooks = make(map[EventID]*Hooks)
	l.active = true

	l.Register(stopListener, func(*Listener, EventData) {})

	go l.listen()

	return
}

func (l *Listener) Stop() (err error) {
	if !l.active {
		err = errors.New("The listener is already inactive!")
		return
	}

	l.active = false

	l.Emit(Event{Event: stopListener, Data: EventData{}})

	return
}

func (l *Listener) ListenOn(in chan Event) {
	l.incoming = in
	l.Stop()
	go l.listen()
}

func (l *Listener) Register(event EventID, callback Callback) {
	if !l.active {
		l.Start(1)
	}

	if _, ok := l.hooks[event]; !ok {
		l.hooks[event] = new(Hooks)
		l.hooks[event].Initialize()
	}

	l.callbacks[event] = callback
}

func (l *Listener) RegisterIfNil(event EventID, callback Callback) (err error) {
	if l.callbacks[event] == nil {
		l.Register(event, callback)
		return
	}

	return errors.New("This event is already set!")
}

func (l *Listener) Registered(event EventID) bool {
	if l.callbacks[event] == nil {
		return false
	}
	return true
}

func (l *Listener) RegisterHook(t int, event EventID, callback Callback) {
	if t == BEFORE {
		l.hooks[event].Before = append(l.hooks[event].Before, callback)
	} else if t == AFTER {
		l.hooks[event].After = append(l.hooks[event].After, callback)
	}
}

func (l *Listener) Emit(event Event) {
	l.incoming <- event
}

func (l *Listener) listen() {
	for {
		e := <-l.incoming

		if e.Event == stopListener {
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

func (l *Listener) runBeforeHooks(e Event) {
	l.runHooks(l.hooks[e.Event].Before, e)
}

func (l *Listener) runAfterHooks(e Event) {
	l.runHooks(l.hooks[e.Event].After, e)
}

func (l *Listener) runHooks(hooks []Callback, e Event) {
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
