package plugin

// EventType is an enumeration of the possible types of events.
type EventType int

// Type is an enumeration of the possible types of plugins.
type Type int

const (
	CollectorFlush EventType = iota
)

const (
	Source Type = iota
	Sink
)

type Event struct {
	Type EventType
	Data interface{}
}
