package eventa

type EventID int

type EventData []interface{}

type Event struct {
	Event      EventID
	Data       EventData
	Concurrent bool
}
