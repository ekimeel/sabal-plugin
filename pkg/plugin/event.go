package plugin

import "github.com/ekimeel/sabal-pb/pb"

// Event is a struct that represents an event that a Plugin can process.
// It includes a slice of Metrics, where each Metric represents a single
// data point or measurement at a certain point in time. The exact structure
// and interpretation of a Metric is dependent on the specific application
// and the specific plugin that processes the event.
type Event struct {
	Metrics []pb.Metric
}
