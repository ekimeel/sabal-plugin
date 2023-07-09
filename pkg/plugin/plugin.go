package plugin

import (
	"github.com/ekimeel/sabal-pb/pb"
	"plugin"
)

type Plugin interface {
	Process(metrics []pb.Metric) error
	Install(env *Environment) error
	Name() string
	Version() string
}

type PluginManager interface {
	RunChannel() chan plugin.Plugin
}
