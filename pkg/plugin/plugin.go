package plugin

import (
	"github.com/ekimeel/sabal-pb/pb"
	"time"
)

type Status int8

const (
	None Status = iota
	Running
	Ready
	Error
)

type Plugin interface {
	Run(offset Offset) error
	Install(env *Environment) error
	Name() string
	Version() string
	Status() Status
	LastRuntime() time.Time
	Offset() Offset
}

type Offset struct {
	Value time.Time
}

type Environment struct {
	Config        map[string]interface{}
	EquipService  pb.EquipServiceClient
	PointService  pb.PointServiceClient
	MetricService pb.MetricServiceClient
}
