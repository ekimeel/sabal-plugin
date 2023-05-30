package plugin

import (
	"plugin"
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
	Run(env *Context) error
	Install(env *Context) error
	Name() string
	Version() string
	Status() Status
	LastRuntime() time.Time
	Offset() Offset
}

type Offset struct {
	Value time.Time
}

type PluginManager interface {
	RunChannel() chan plugin.Plugin
}
