package entity

type EventType string

const (
	EventTypeCreate EventType = "CREATE"
	EventTypeUpdate EventType = "UPDATE"
	EventTypeDelete EventType = "DELETE"
)

type Event struct {
	Key  string
	Type EventType
}
