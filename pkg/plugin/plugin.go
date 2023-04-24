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
	Config() Config
	Status() int8
	LastRuntime() time.Time
}

type Offset struct {
	Value time.Time
}

type Environment struct {
	EquipService  pb.EquipServiceClient
	PointService  pb.PointServiceClient
	MetricService pb.MetricServiceClient
}
