package plugin

// EventType is an enumeration of the possible types of events.
type EventType int

const (
	CollectorFlush EventType = iota
)

type Event struct {
	Type EventType
	Data interface{}
}
