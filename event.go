package eventa

type Event struct {
	Event      int
	Data       map[string]interface{}
	Concurrent bool
}

func NewEvent(eventID int) *Event {
	e := &Event{Event: eventID}
	e.Data = make(map[string]interface{})
	return e
}
