package plugin

import "github.com/ekimeel/sabal-pb/pb"

// EventType is an enumeration of the possible types of events.
type EventType int

const (
	// Pipeline represents a pipeline event.
	Pipeline EventType = iota
	// Maintenance represents a maintenance event.
	Maintenance
	// Startup represents a startup event.
	Startup
	// Shutdown represents a shutdown event.
	Shutdown
)

// Event is a struct that represents an event that a Plugin can process.
// It includes a slice of Metrics, where each Metric represents a single
// data point or measurement at a certain point in time. The exact structure
// and interpretation of a Metric is dependent on the specific application
// and the specific plugin that processes the event.
type Event struct {
	Type    EventType
	Metrics []pb.Metric
}
