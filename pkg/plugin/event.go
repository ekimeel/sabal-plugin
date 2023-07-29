package plugin

// EventType is an enumeration of the possible types of events.
type EventType int

// PluginType is an enumeration of the possible types of plugins.
type PluginType int

const (
	CollectorFlush EventType = iota
)

const (
	Source PluginType = iota
	Sink
)

type Event struct {
	Type EventType
	Data interface{}
}
